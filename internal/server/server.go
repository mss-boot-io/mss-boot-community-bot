package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/mss-boot-io/mss-boot-community-bot/internal/config"
	"github.com/mss-boot-io/mss-boot-community-bot/internal/policy"
)

const maxWebhookBodyBytes = 1 << 20

type Server struct {
	cfg    config.Config
	logger *slog.Logger
	mux    *http.ServeMux
}

func New(cfg config.Config, logger *slog.Logger) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}

	s := &Server{
		cfg:    cfg,
		logger: logger,
		mux:    http.NewServeMux(),
	}
	s.routes()
	return s.mux
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.healthz)
	s.mux.HandleFunc("GET /readyz", s.readyz)
	s.mux.HandleFunc("POST /webhooks/github", s.githubWebhook)
	s.mux.HandleFunc("POST /webhooks/qq", s.qqWebhook)
	s.mux.HandleFunc("POST /policy/evaluate", s.evaluatePolicy)
}

func (s *Server) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) readyz(w http.ResponseWriter, _ *http.Request) {
	integrations := map[string]string{
		"github_webhook": "optional",
		"llm":            "missing",
		"qq_bot":         "missing",
		"wechat":         "manual-or-open-api",
	}
	if s.cfg.GitHubWebhookSecret != "" {
		integrations["github_webhook"] = "signature-verification-enabled"
	}
	if s.cfg.LLM.Configured() {
		integrations["llm"] = "configured"
	}
	if s.cfg.QQAppID != "" && s.cfg.QQAppSecret != "" {
		integrations["qq_bot"] = "configured"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"status":       "ready",
		"integrations": integrations,
	})
}

func (s *Server) githubWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if s.cfg.GitHubWebhookSecret != "" && !validGitHubSignature(r.Header.Get("X-Hub-Signature-256"), body, s.cfg.GitHubWebhookSecret) {
		writeError(w, http.StatusUnauthorized, errors.New("invalid github signature"))
		return
	}

	event := r.Header.Get("X-GitHub-Event")
	delivery := r.Header.Get("X-GitHub-Delivery")
	s.logger.Info("accepted github webhook", "event", event, "delivery", delivery, "bytes", len(body))
	writeJSON(w, http.StatusAccepted, map[string]string{"status": "accepted"})
}

func (s *Server) qqWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	s.logger.Info("accepted qq webhook placeholder", "bytes", len(body))
	writeJSON(w, http.StatusAccepted, map[string]string{
		"status": "accepted",
		"note":   "official QQ bot verification and event parsing are implemented in a follow-up adapter PR",
	})
}

func (s *Server) evaluatePolicy(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var request struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal(body, &request); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, policy.EvaluateMessage(request.Text))
}

func readBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(http.MaxBytesReader(nil, r.Body, maxWebhookBodyBytes))
}

func validGitHubSignature(header string, body []byte, secret string) bool {
	signature, ok := strings.CutPrefix(header, "sha256=")
	if !ok {
		return false
	}

	provided, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	return hmac.Equal(provided, mac.Sum(nil))
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}
