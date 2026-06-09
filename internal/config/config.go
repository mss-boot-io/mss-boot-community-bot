package config

import "os"

type Config struct {
	Addr                string
	GitHubWebhookSecret string
	QQAppID             string
	QQAppSecret         string
}

func LoadFromEnv() Config {
	return Config{
		Addr:                getenv("MSS_BOT_ADDR", ":8080"),
		GitHubWebhookSecret: os.Getenv("MSS_BOT_GITHUB_WEBHOOK_SECRET"),
		QQAppID:             os.Getenv("MSS_BOT_QQ_APP_ID"),
		QQAppSecret:         os.Getenv("MSS_BOT_QQ_APP_SECRET"),
	}
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
