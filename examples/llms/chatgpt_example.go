package main

import (
	"context"
	"fmt"
	"log"

	chatgpt "github.com/index-labs/pipelang/pkg/llms/chatgpt"
)

func main() {
	ctx := context.Background()
	llm, err := chatgpt.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	messages := make([]chatgpt.Message, 0)
	messages = append(messages, chatgpt.SystemMessage("You are a helpful assistant that translates English to French."))
	messages = append(messages, chatgpt.HumanMessage("Translate this sentence from English to French. I love programming."))

	result, err := llm.Call(ctx, messages...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
