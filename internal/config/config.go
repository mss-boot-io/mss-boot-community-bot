package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Addr                string
	GitHubWebhookSecret string
	QQAppID             string
	QQAppSecret         string
	LLM                 LLMConfig
}

type LLMConfig struct {
	BaseURL     string
	APIKey      string
	Model       string
	Timeout     time.Duration
	Temperature float64
	MaxTokens   int
}

func (c LLMConfig) Configured() bool {
	return c.BaseURL != "" && c.APIKey != "" && c.Model != ""
}

func LoadFromEnv() Config {
	return Config{
		Addr:                getenv("MSS_BOT_ADDR", ":8080"),
		GitHubWebhookSecret: os.Getenv("MSS_BOT_GITHUB_WEBHOOK_SECRET"),
		QQAppID:             os.Getenv("MSS_BOT_QQ_APP_ID"),
		QQAppSecret:         os.Getenv("MSS_BOT_QQ_APP_SECRET"),
		LLM: LLMConfig{
			BaseURL:     getenv("MSS_BOT_LLM_BASE_URL", "https://api.openai.com/v1"),
			APIKey:      os.Getenv("MSS_BOT_LLM_API_KEY"),
			Model:       os.Getenv("MSS_BOT_LLM_MODEL"),
			Timeout:     getDuration("MSS_BOT_LLM_TIMEOUT", 30*time.Second),
			Temperature: getFloat("MSS_BOT_LLM_TEMPERATURE", 0.2),
			MaxTokens:   getInt("MSS_BOT_LLM_MAX_TOKENS", 1024),
		},
	}
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func getFloat(key string, fallback float64) float64 {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			return parsed
		}
	}
	return fallback
}

func getInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}
