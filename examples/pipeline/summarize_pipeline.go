package main

import (
	"context"
	"fmt"
	"os"

	"github.com/index-labs/pipelang/pkg/constants"
	"github.com/index-labs/pipelang/pkg/llms/openai"
	"github.com/index-labs/pipelang/pkg/loaders"
	"github.com/index-labs/pipelang/pkg/pipeline/summarize"
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
