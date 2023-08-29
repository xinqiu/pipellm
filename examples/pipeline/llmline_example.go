package main

import (
	"context"
	"fmt"
	"github.com/xinqiu/pipellm/pkg/prompts"

	"github.com/xinqiu/pipellm/pkg/llms/openai"
	"github.com/xinqiu/pipellm/pkg/pipeline"
)

func main() {
	ctx := context.Background()
	prompt := prompts.New(ctx, []string{"product"}, "What is a good name for a company that makes {{.product}}?")
	llm, _ := openai.New(ctx)
	llmLine := pipeline.New(ctx, prompt, llm)
	result, _ := llmLine.Run(ctx, "colorful socks")
	fmt.Print(result)
}
