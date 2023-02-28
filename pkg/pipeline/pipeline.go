package pipeline

import "context"

// BasePipeline represents base pipeline interface
type BasePipeline interface {
	Name() string
}

// Pipeline represents pipeline interface
type Pipeline interface {
	BasePipeline
	Run(ctx context.Context)
}
