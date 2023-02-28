package openai

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/index-labs/pipelang/pkg/llms"

	constantOpenAI "github.com/index-labs/pipelang/pkg/constants"
	"github.com/index-labs/pipelang/pkg/types"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// LLM OpenAI implementation
type LLM struct {
	// ========== LLM Client ===========
	client *gogpt.Client

	// ========== Parameters =========
	// modelName Model name to use.
	modelName string
	// temperature What sampling temperature to use.
	temperature float32
	// maxTokens The maximum number of tokens to generate in the completion.
	// 	-1 returns as many tokens as possible given the prompt and the models maximal context size.
	maxTokens int
	// topP Total probability mass of tokens to consider at each step.
	topP float32
	// frequencyPenalty Penalizes repeated tokens according to frequency.
	frequencyPenalty float32
	// presencePenalty Penalizes repeated tokens.
	presencePenalty float32
	// n How many completions to generate for each prompt.
	n int
	// bestOf Generates best_of completions server-side and returns the "best".
	bestOf int
	// modelKwargs Holds any model parameters valid for `create` call not explicitly specified.
	modelKwargs map[string]interface{}
	// openaiApiKey
	openaiAPIKey string
	// batchSize Batch size to use when passing multiple documents to generate.
	batchSize int
	// requestTimeout Timeout for requests to OpenAI completion API. Default is 600 seconds.
	requestTimeout float64
	// logitBias Adjust the probability of specific tokens being generated.
	logitBias map[string]float64
	// maxRetries Maximum number of retries to make when generating.
	maxRetries int
	// streaming Whether to stream the results or not.
	streaming bool
}

// ModelName is the name of the model
func (receiver *LLM) ModelName() string {
	return "OpenAI"
}

// Call is the method to call OpenAI
func (receiver *LLM) Call(ctx context.Context, prompt string) (string, error) {
	result, err := receiver.Generate(ctx, []string{prompt})
	if err != nil {
		return "", err
	}
	fmt.Printf("token_usage:%+v", result.LLMOutput.(*gogpt.CompletionResponse).Usage)

	return result.Generations[0].Text, nil
}

// Generate OpenAI implementation
func (receiver *LLM) Generate(ctx context.Context, prompts []string) (*llms.LLMResult, error) {
	if receiver.streaming {
		// TODO: not support streaming now
		return nil, types.ErrNotImplemented
	}

	completionRequest := gogpt.CompletionRequest{
		Model:       receiver.modelName,
		Prompt:      strings.Join(prompts, "\n"),
		MaxTokens:   receiver.maxTokens,
		Temperature: receiver.temperature,
		TopP:        receiver.topP,
		N:           receiver.n,
		BestOf:      receiver.bestOf,
	}
	resp, err := receiver.client.CreateCompletion(ctx, completionRequest)
	if err != nil {
		return nil, err
	}
	if len(resp.Choices) == 0 {
		return nil, types.ErrOpenAIResponse
	}

	return receiver.createLLMResult(ctx, &resp), nil
}

func (receiver *LLM) createLLMResult(ctx context.Context, response *gogpt.CompletionResponse) *llms.LLMResult {
	generatons := make([]*llms.Generation, 0)

	for _, choice := range response.Choices {
		generatons = append(generatons, &llms.Generation{
			Text: choice.Text,
			GenerationInfo: map[string]any{
				"finish_reason": choice.FinishReason,
				"logprobs":      choice.LogProbs,
			},
		})
	}

	return &llms.LLMResult{
		Generations: generatons,
		LLMOutput:   response,
	}
}

// New return a new LLM client
func New(ctx context.Context) (*LLM, error) {
	apiKey := os.Getenv(constantOpenAI.OsEnvAPIKey)
	if apiKey == "" {
		return nil, types.ErrNoAPIKey
	}
	client := gogpt.NewClient(apiKey)
	llm := &LLM{client: client, openaiAPIKey: apiKey}

	// init default parameters
	llm.buildDefaultModelParams(ctx)
	return llm, nil
}

// ======= Parameters Builder ========

func (receiver *LLM) buildDefaultModelParams(ctx context.Context) {
	receiver.modelName = gogpt.GPT3TextDavinci003
	receiver.temperature = 0.7
	receiver.maxTokens = 256
	receiver.topP = 1
	receiver.frequencyPenalty = 0
	receiver.presencePenalty = 0
	receiver.n = 1
	receiver.bestOf = 1
	receiver.batchSize = 20
	receiver.maxRetries = 6
	receiver.streaming = false
}

// WithModelName add model name to receiver
func (receiver *LLM) WithModelName(modelName string) *LLM {
	receiver.modelName = modelName
	return receiver
}

// WithTemperature add temperature to receiver
func (receiver *LLM) WithTemperature(temperature float32) *LLM {
	receiver.temperature = temperature
	return receiver
}

// WithMaxTokens add maxTokens to receiver
func (receiver *LLM) WithMaxTokens(maxTokens int) *LLM {
	receiver.maxTokens = maxTokens
	return receiver
}
