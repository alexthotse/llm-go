package openai

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/simonw/llm-go/pkg/llm"
)

func TestOpenAIModelExecute(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.openai.com/v1/chat/completions",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{
				"id": "chatcmpl-123",
				"object": "chat.completion",
				"created": 1677652288,
				"model": "gpt-3.5-turbo-0613",
				"choices": [{
					"index": 0,
					"message": {
						"role": "assistant",
						"content": "\n\nHello there, how may I assist you today?"
					},
					"finish_reason": "stop"
				}],
				"usage": {
					"prompt_tokens": 9,
					"completion_tokens": 12,
					"total_tokens": 21
				}
			}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	model := NewOpenAIModel("gpt-3.5-turbo")
	prompt := &llm.Prompt{
		Prompt: "hello",
	}
	response, err := model.Execute(prompt, nil)
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if response == nil {
		t.Fatal("response is nil")
	}
	expected := "\n\nHello there, how may I assist you today?"
	if response.Text != expected {
		t.Errorf("response.Text = %q, want %q", response.Text, expected)
	}

	// Verify that a request was made
	if httpmock.GetTotalCallCount() != 1 {
		t.Errorf("expected 1 request, got %d", httpmock.GetTotalCallCount())
	}
}

func TestOpenAIModelExecuteError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.openai.com/v1/chat/completions",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(401, `{
				"error": {
					"message": "Incorrect API key provided: sk-.... You can find your API key at https://platform.openai.com/account/api-keys.",
					"type": "invalid_request_error",
					"param": null,
					"code": "invalid_api_key"
				}
			}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	model := NewOpenAIModel("gpt-3.5-turbo")
	prompt := &llm.Prompt{
		Prompt: "hello",
	}
	_, err := model.Execute(prompt, nil)
	if err == nil {
		t.Fatal("Execute() error is nil, want error")
	}
	expected := "Incorrect API key provided"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("Execute() error = %q, want to contain %q", err.Error(), expected)
	}
}
