package llms

import "context"

// LLM is the interface implemented by an object which is a LLM
type LLM interface {
	Call(ctx context.Context, prompt string) (string, error)
	Generate(ctx context.Context, prompts []string)
}

type LLMResult struct {
	Generations []*Generation
	LLMOutput   any
}

type Generation struct {
	Text           string         `json:"text"`
	GenerationInfo map[string]any `json:"generationInfo"`
}
