package main

import (
	"context"
	"fmt"

	"github.com/xinqiu/pipellm/pkg/prompts"
)

func main() {
	ctx := context.Background()
	prompt := prompts.New(ctx, []string{"adjective", "content"}, "Tell me a {{.adjective}} joke about {{.content}}.")
	fmt.Print(prompt.Format(ctx, map[string]string{
		"adjective": "funny",
		"content":   "chickens",
	}))
}
