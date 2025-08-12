package openai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/simonw/llm-go/pkg/llm"
)

// OpenAIModel is a model that uses the OpenAI API.
type OpenAIModel struct {
	name   string
	client *openai.Client
}

// NewOpenAIModel creates a new OpenAIModel.
func NewOpenAIModel(name string) *OpenAIModel {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	return &OpenAIModel{
		name:   name,
		client: client,
	}
}

// Name returns the name of the model.
func (m *OpenAIModel) Name() string {
	return m.name
}

// Execute executes a prompt against the model.
func (m *OpenAIModel) Execute(prompt *llm.Prompt, options *llm.Options) (*llm.Response, error) {
	resp, err := m.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: m.name,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt.Prompt,
				},
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("error from OpenAI API: %w", err)
	}

	return &llm.Response{
		Text: resp.Choices[0].Message.Content,
	}, nil
}
