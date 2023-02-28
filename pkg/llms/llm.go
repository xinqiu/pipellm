package llms

import "context"

// BaseLLM is base LLM interface
type BaseLLM interface {
	ModelName() string
}

// LLM is the interface implemented by an object which is a LLM
type LLM interface {
	BaseLLM
	Call(ctx context.Context, prompt string) (string, error)
	Generate(ctx context.Context, prompts []string) (*LLMResult, error)
}

// LLMResult is result returned from LLM
type LLMResult struct {
	Generations []*Generation
	LLMOutput   any
}

// Generation is LLM output
type Generation struct {
	Text           string         `json:"text"`
	GenerationInfo map[string]any `json:"generationInfo"`
}
