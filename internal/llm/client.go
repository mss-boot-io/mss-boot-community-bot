package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Options struct {
	BaseURL     string
	APIKey      string
	Model       string
	Timeout     time.Duration
	Temperature float64
	MaxTokens   int
}

type Client struct {
	endpoint    string
	apiKey      string
	model       string
	temperature float64
	maxTokens   int
	httpClient  *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

type chatCompletionResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func NewClient(opts Options) (*Client, error) {
	baseURL := strings.TrimRight(strings.TrimSpace(opts.BaseURL), "/")
	if baseURL == "" {
		return nil, errors.New("llm base url is required")
	}
	if opts.APIKey == "" {
		return nil, errors.New("llm api key is required")
	}
	if opts.Model == "" {
		return nil, errors.New("llm model is required")
	}

	timeout := opts.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	return &Client{
		endpoint:    baseURL + "/chat/completions",
		apiKey:      opts.APIKey,
		model:       opts.Model,
		temperature: opts.Temperature,
		maxTokens:   opts.MaxTokens,
		httpClient:  &http.Client{Timeout: timeout},
	}, nil
}

func (c *Client) Complete(ctx context.Context, messages []Message) (string, error) {
	if len(messages) == 0 {
		return "", errors.New("at least one message is required")
	}

	payload := chatCompletionRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: c.temperature,
		MaxTokens:   c.maxTokens,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		errorBody, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return "", fmt.Errorf("llm request failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(errorBody)))
	}

	var decoded chatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return "", err
	}
	if len(decoded.Choices) == 0 {
		return "", errors.New("llm response contains no choices")
	}
	return decoded.Choices[0].Message.Content, nil
}
