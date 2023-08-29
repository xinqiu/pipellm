package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xinqiu/pipellm/pkg/loaders"
)

func main() {
	ctx := context.Background()
	currentDirectory, _ := os.Getwd()
	doc, _ := loaders.New(ctx).Load(ctx, currentDirectory+"/examples/loaders/demo.txt")
	fmt.Println(doc[0].PageContent, doc[0].Metadata)
}
