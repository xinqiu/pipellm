package pipeline

import (
	"context"

	"github.com/index-labs/pipelang/pkg/llms"
	"github.com/index-labs/pipelang/pkg/prompts"
	"github.com/index-labs/pipelang/pkg/types"
	"github.com/index-labs/pipelang/pkg/utils"
)

// LLMLine pipeline implementation
type LLMLine struct {
	Pipeline
	prompt prompts.Prompt
	llm    llms.LLM
}

// Name returns the name
func (receiver *LLMLine) Name() string {
	return "LLMLine"
}

// Run runs the LLMLine pipeline
func (receiver *LLMLine) Run(ctx context.Context, inputArgs string) (string, error) {
	if !promptInputVarsValid(ctx, receiver.prompt) {
		return "", types.ErrOnlySupportOneArg
	}

	input := utils.MergeKVListToMap(receiver.prompt.InputVars(ctx), []string{inputArgs})
	result, err := receiver.llm.Call(ctx, receiver.prompt.Format(ctx, input))
	if err != nil {
		return "", err
	}
	return result, nil
}

// New create a LLMLine
func New(ctx context.Context, prompt prompts.Prompt, llm llms.LLM) *LLMLine {
	return &LLMLine{
		prompt: prompt,
		llm:    llm,
	}
}

func promptInputVarsValid(ctx context.Context, prompt prompts.Prompt) bool {
	if len(prompt.InputVars(ctx)) == 1 {
		return true
	}
	return false
}
