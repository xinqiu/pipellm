package main

import (
	"context"
	"fmt"

	chatllm "github.com/xinqiu/pipellm/pkg/llms/chatgpt"
	"github.com/xinqiu/pipellm/pkg/prompts/chatgpt"
)

func main() {
	ctx := context.Background()
	prompt := chatgpt.New(ctx, chatllm.ChatMessageRoleUser, []string{"adjective", "content"}, "Tell me a {{.adjective}} joke about {{.content}}.")
	fmt.Print(prompt.Format(ctx, map[string]string{
		"adjective": "funny",
		"content":   "chickens",
	}))
}
