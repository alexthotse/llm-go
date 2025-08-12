package llm

// Model is the interface for a language model.
type Model interface {
	Name() string
	Execute(prompt *Prompt, options *Options) (*Response, error)
}

// Prompt is a prompt to be sent to a language model.
type Prompt struct {
	System string
	Prompt string
	// ... and other fields
}

// Options are model-specific options.
type Options struct {
	// ... model-specific options
}

// Response is a response from a language model.
type Response struct {
	ID           string
	Prompt       *Prompt
	Text         string
	Model        Model
	// ... and other fields
}
