package loaders

import (
	"context"

	"github.com/xinqiu/pipellm/pkg/types"
)

// BaseLoader represents base loader interface
type BaseLoader interface {
	Name() string
}

// Loader represents loader interface
type Loader interface {
	Load(ctx context.Context, filename string) ([]*types.Document, error)
}
