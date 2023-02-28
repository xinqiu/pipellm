package prompts

import "context"

// BasePrompt base prompt interface
type BasePrompt interface {
	Name() string
}

// Prompt prompt interface
type Prompt interface {
	BasePrompt
	Format(ctx context.Context, inputVars map[string]string) string
}
