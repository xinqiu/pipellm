package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xinqiu/pipellm/pkg/constants"
	"github.com/xinqiu/pipellm/pkg/llms/openai"
	"github.com/xinqiu/pipellm/pkg/loaders"
	"github.com/xinqiu/pipellm/pkg/pipeline/summarize"
)

func main() {
	ctx := context.Background()
	llm, _ := openai.New(ctx)
	currentDirectory, _ := os.Getwd()
	doc, _ := loaders.New(ctx).Load(ctx, currentDirectory+"/examples/loaders/demo.txt")
	doc[0].PageContent = doc[0].PageContent[:300]
	summarizePipeline := summarize.LoadSummarizePipeline(ctx, llm, constants.ChainTypeMapReduce)
	result, _ := summarizePipeline.Run(ctx, doc)
	fmt.Print(result)
}
