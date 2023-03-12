package chatgpt

import (
	"context"
	"fmt"
	"os"

	constantOpenAI "github.com/index-labs/pipelang/pkg/constants"
	"github.com/index-labs/pipelang/pkg/llms"
	"github.com/index-labs/pipelang/pkg/types"
	openai "github.com/sashabaranov/go-openai"
)

// LLMChat chat model interface
type LLMChat interface {
	llms.BaseLLM
	Call(ctx context.Context, message ...Message) (string, error)
	Generate(ctx context.Context, message []Message) (*llms.LLMResult, error)
}

// ChatLLM OpenAI chatgpt implementation
type ChatLLM struct {
	// ========== ChatLLM Client ===========
	client *openai.Client

	// ========== Parameters =========
	// modelName Model name to use.
	modelName string
	// openaiApiKey
	openaiAPIKey string
}

// ModelName is the name of the model
func (receiver *ChatLLM) ModelName() string {
	return "OpenAI ChatGPT"
}

// Message type of message
type Message openai.ChatCompletionMessage

// Call is the method to call OpenAI
func (receiver *ChatLLM) Call(ctx context.Context, message ...Message) (string, error) {
	result, err := receiver.Generate(ctx, message)
	if err != nil {
		return "", err
	}
	fmt.Printf("token_usage:%+v", result.LLMOutput.(*openai.ChatCompletionResponse).Usage)

	return result.Generations[0].Text, nil
}

// Generate OpenAI implementation
func (receiver *ChatLLM) Generate(ctx context.Context, message []Message) (*llms.LLMResult, error) {
	chatRequest := openai.ChatCompletionRequest{
		Model:    receiver.modelName,
		Messages: converChatGPTMessageTypeToOpenAIChatMessageType(message),
	}

	resp, err := receiver.client.CreateChatCompletion(ctx, chatRequest)
	if err != nil {
		return nil, err
	}
	if len(resp.Choices) == 0 {
		return nil, types.ErrOpenAIResponse
	}

	return receiver.createChatResult(ctx, &resp), nil
}

// New return a new ChatLLM client
func New(ctx context.Context) (*ChatLLM, error) {
	apiKey := os.Getenv(constantOpenAI.OsEnvAPIKey)
	if apiKey == "" {
		return nil, types.ErrNoAPIKey
	}
	client := openai.NewClient(apiKey)
	llm := &ChatLLM{client: client, openaiAPIKey: apiKey}

	// init default parameters
	llm.buildDefaultModelParams(ctx)
	return llm, nil
}

func (receiver *ChatLLM) createChatResult(ctx context.Context, response *openai.ChatCompletionResponse) *llms.LLMResult {
	generatons := make([]*llms.Generation, 0)

	for _, choice := range response.Choices {
		generatons = append(generatons, &llms.Generation{
			Text: choice.Message.Content,
			GenerationInfo: map[string]any{
				"finish_reason": choice.FinishReason,
			},
		})
	}

	return &llms.LLMResult{
		Generations: generatons,
		LLMOutput:   response,
	}
}

// ======= Parameters Builder ========

func (receiver *ChatLLM) buildDefaultModelParams(ctx context.Context) {
	receiver.modelName = openai.GPT3Dot5Turbo
}

// WithModelName add model name to receiver
func (receiver *ChatLLM) WithModelName(modelName string) *ChatLLM {
	receiver.modelName = modelName
	return receiver
}

// HumanMessage human message of ChatGPT message
func HumanMessage(content string) Message {
	return Message{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	}
}

// SystemMessage system message of ChatGPT message
func SystemMessage(content string) Message {
	return Message{
		Role:    openai.ChatMessageRoleSystem,
		Content: content,
	}
}

// AssistantMessage assistant message of ChatGPT message
func AssistantMessage(content string) Message {
	return Message{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	}
}

func converChatGPTMessageTypeToOpenAIChatMessageType(messages []Message) []openai.ChatCompletionMessage {
	msgs := make([]openai.ChatCompletionMessage, 0)
	for _, message := range messages {
		msgs = append(msgs, openai.ChatCompletionMessage(message))
	}
	return msgs
}
