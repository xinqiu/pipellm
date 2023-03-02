package summarize

import (
	"context"

	"github.com/index-labs/pipelang/pkg/constants"
	"github.com/index-labs/pipelang/pkg/llms"
	"github.com/index-labs/pipelang/pkg/types"
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
