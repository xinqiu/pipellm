package summarize

import (
	"context"

	"github.com/index-labs/pipelang/pkg/constants"
	"github.com/index-labs/pipelang/pkg/llms"
	"github.com/index-labs/pipelang/pkg/prompts"
	"github.com/index-labs/pipelang/pkg/types"
	"github.com/index-labs/pipelang/pkg/utils"
)

// MapReduceLine pipeline implementation
type MapReduceLine struct {
	prompt *prompts.PromptTemplate
	llm    llms.LLM
}

// Name returns the name
func (receiver *MapReduceLine) Name() string {
	return "MapReduceLine"
}

// New returns a new MapReduceLine
func New(ctx context.Context, llm llms.LLM) interface{} {
	return &MapReduceLine{
		prompt: prompts.New(ctx, constants.SummarizeMapReducePromptInputVars, constants.SummarizeMapReducePromptTemplate),
		llm:    llm,
	}
}

// Run run pipeline implementation
func (receiver *MapReduceLine) Run(ctx context.Context, document []*types.Document) (string, error) {
	input := utils.MergeKVListToMap(receiver.prompt.InputVars(ctx), []string{document[0].PageContent})
	result, err := receiver.llm.Call(ctx, receiver.prompt.Format(ctx, input))
	if err != nil {
		return "", err
	}
	return result, nil
}
