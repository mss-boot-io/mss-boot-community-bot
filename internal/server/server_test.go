package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mss-boot-io/mss-boot-community-bot/internal/config"
)

func TestHealthz(t *testing.T) {
	handler := New(config.Config{}, nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestGitHubWebhookSignature(t *testing.T) {
	body := []byte(`{"zen":"Keep it logically awesome."}`)
	secret := "test-secret"
	handler := New(config.Config{GitHubWebhookSecret: secret}, nil)

	req := httptest.NewRequest(http.MethodPost, "/webhooks/github", bytes.NewReader(body))
	req.Header.Set("X-Hub-Signature-256", "sha256="+sign(body, secret))
	req.Header.Set("X-GitHub-Event", "ping")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want %d: %s", rec.Code, http.StatusAccepted, rec.Body.String())
	}
}

func TestGitHubWebhookRejectsInvalidSignature(t *testing.T) {
	handler := New(config.Config{GitHubWebhookSecret: "test-secret"}, nil)
	req := httptest.NewRequest(http.MethodPost, "/webhooks/github", bytes.NewReader([]byte("{}")))
	req.Header.Set("X-Hub-Signature-256", "sha256=bad")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func sign(body []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}
