package llm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientCompleteUsesOpenAICompatibleChatCompletions(t *testing.T) {
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Fatalf("authorization = %q", got)
		}

		var request chatCompletionRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal(err)
		}
		if request.Model != "test-model" {
			t.Fatalf("model = %q", request.Model)
		}
		if len(request.Messages) != 1 || request.Messages[0].Content != "hello" {
			t.Fatalf("messages = %#v", request.Messages)
		}

		_ = json.NewEncoder(w).Encode(map[string]any{
			"choices": []map[string]any{
				{"message": map[string]string{"role": "assistant", "content": "world"}},
			},
		})
	}))
	defer api.Close()

	client, err := NewClient(Options{
		BaseURL: api.URL + "/v1",
		APIKey:  "test-key",
		Model:   "test-model",
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Complete(context.Background(), []Message{{Role: "user", Content: "hello"}})
	if err != nil {
		t.Fatal(err)
	}
	if got != "world" {
		t.Fatalf("completion = %q", got)
	}
}

func TestNewClientRequiresModel(t *testing.T) {
	_, err := NewClient(Options{BaseURL: "https://llm.example.com/v1", APIKey: "test-key"})
	if err == nil {
		t.Fatal("expected error")
	}
}
