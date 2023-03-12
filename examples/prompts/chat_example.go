package main

import (
	"context"
	"fmt"

	chatllm "github.com/index-labs/pipelang/pkg/llms/chatgpt"
	"github.com/index-labs/pipelang/pkg/prompts/chatgpt"
)

func main() {
	ctx := context.Background()
	prompt := chatgpt.New(ctx, chatllm.ChatMessageRoleUser, []string{"adjective", "content"}, "Tell me a {{.adjective}} joke about {{.content}}.")
	fmt.Print(prompt.Format(ctx, map[string]string{
		"adjective": "funny",
		"content":   "chickens",
	}))
}
