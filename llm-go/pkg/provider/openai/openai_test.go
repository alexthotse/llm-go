package openai

import (
	"testing"

	"github.com/simonw/llm-go/pkg/llm"
)

func TestOpenAIModelExecute(t *testing.T) {
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
	expected := "This is a dummy response from the gpt-3.5-turbo model."
	if response.Text != expected {
		t.Errorf("response.Text = %q, want %q", response.Text, expected)
	}
}
