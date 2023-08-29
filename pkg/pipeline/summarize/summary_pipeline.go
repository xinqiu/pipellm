package summarize

import (
	"context"

	"github.com/xinqiu/pipellm/pkg/constants"
	"github.com/xinqiu/pipellm/pkg/llms"
	"github.com/xinqiu/pipellm/pkg/types"
)

// SummaryLine interface
type SummaryLine interface {
	Run(ctx context.Context, document []*types.Document) (string, error)
}

// LoadSummarizePipeline chooses a pipeline by chain type
func LoadSummarizePipeline(ctx context.Context, llm llms.LLM, chainType constants.ChainType) SummaryLine {
	switch chainType {
	case constants.ChainTypeMapReduce:
		return New(ctx, llm).(SummaryLine)
	}
	return nil
}
