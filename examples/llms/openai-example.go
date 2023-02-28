package main

import (
	"context"
	"fmt"
	"log"

	gogpt "github.com/sashabaranov/go-gpt3"

	"github.com/index-labs/pipelang/pkg/llms/openai"
)

func main() {
	ctx := context.Background()
	llm, err := openai.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	llm.WithModelName(gogpt.GPT3TextDavinci002).WithTemperature(0.5)

	result, err := llm.Call(ctx, "What would be a good company name a company that makes colorful socks?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
