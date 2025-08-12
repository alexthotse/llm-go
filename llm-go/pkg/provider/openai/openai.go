package openai

import (
	"fmt"

	"github.com/simonw/llm-go/pkg/llm"
)

// OpenAIModel is a model that uses the OpenAI API.
type OpenAIModel struct {
	name string
}

// NewOpenAIModel creates a new OpenAIModel.
func NewOpenAIModel(name string) *OpenAIModel {
	return &OpenAIModel{name: name}
}

// Name returns the name of the model.
func (m *OpenAIModel) Name() string {
	return m.name
}

// Execute executes a prompt against the model.
func (m *OpenAIModel) Execute(prompt *llm.Prompt, options *llm.Options) (*llm.Response, error) {
	// For now, just return a dummy response.
	return &llm.Response{
		Text: fmt.Sprintf("This is a dummy response from the %s model.", m.name),
	}, nil
}
