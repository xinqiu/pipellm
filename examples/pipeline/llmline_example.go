package main

import (
	"context"
	"fmt"

	"github.com/index-labs/pipelang/pkg/prompts/chatgpt"

	"github.com/index-labs/pipelang/pkg/llms/openai"
	"github.com/index-labs/pipelang/pkg/pipeline"
)

func main() {
	ctx := context.Background()
	prompt := chatgpt.New(ctx, []string{"product"}, "What is a good name for a company that makes {{.product}}?")
	llm, _ := openai.New(ctx)
	llmLine := pipeline.New(ctx, prompt, llm)
	result, _ := llmLine.Run(ctx, "colorful socks")
	fmt.Print(result)
}
