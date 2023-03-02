package loaders

import (
	"context"
	"fmt"
	"os"

	"github.com/index-labs/pipelang/pkg/types"
)

// TextLoader loader implementation
type TextLoader struct {
	Loader
}

// Name of loader
func (receiver *TextLoader) Name() string {
	return "TextLoader"
}

// New create a TextLoader
func New(ctx context.Context) *TextLoader {
	return &TextLoader{}
}

// Load implementation
func (receiver *TextLoader) Load(ctx context.Context, filename string) ([]*types.Document, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, types.ErrTextLoader
	}
	docs := make([]*types.Document, 0)
	docs = append(docs, types.NewDocument(string(content), map[string]interface{}{
		"source": filename,
	}))

	return docs, nil
}
