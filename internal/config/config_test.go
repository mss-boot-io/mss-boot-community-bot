package config

import (
	"testing"
	"time"
)

func TestLoadFromEnvLLMConfig(t *testing.T) {
	t.Setenv("MSS_BOT_LLM_BASE_URL", "https://llm.example.com/v1")
	t.Setenv("MSS_BOT_LLM_API_KEY", "test-key")
	t.Setenv("MSS_BOT_LLM_MODEL", "test-model")
	t.Setenv("MSS_BOT_LLM_TIMEOUT", "12s")
	t.Setenv("MSS_BOT_LLM_TEMPERATURE", "0.7")
	t.Setenv("MSS_BOT_LLM_MAX_TOKENS", "2048")

	cfg := LoadFromEnv()

	if !cfg.LLM.Configured() {
		t.Fatal("expected LLM to be configured")
	}
	if cfg.LLM.BaseURL != "https://llm.example.com/v1" {
		t.Fatalf("base url = %q", cfg.LLM.BaseURL)
	}
	if cfg.LLM.Timeout != 12*time.Second {
		t.Fatalf("timeout = %s", cfg.LLM.Timeout)
	}
	if cfg.LLM.Temperature != 0.7 {
		t.Fatalf("temperature = %v", cfg.LLM.Temperature)
	}
	if cfg.LLM.MaxTokens != 2048 {
		t.Fatalf("max tokens = %d", cfg.LLM.MaxTokens)
	}
}
