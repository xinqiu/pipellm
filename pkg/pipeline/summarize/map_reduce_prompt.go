package summarize

import (
	"context"

	"github.com/xinqiu/pipellm/pkg/prompts/chatgpt"

	"github.com/xinqiu/pipellm/pkg/constants"
	"github.com/xinqiu/pipellm/pkg/llms"
	"github.com/xinqiu/pipellm/pkg/prompts"
	"github.com/xinqiu/pipellm/pkg/types"
	"github.com/xinqiu/pipellm/pkg/utils"
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
		prompt: chatgpt.New(ctx, constants.SummarizeMapReducePromptInputVars, constants.SummarizeMapReducePromptTemplate),
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
