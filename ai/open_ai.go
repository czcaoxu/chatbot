package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
	client *openai.Client
}

func NewOpenAIModel(apiKey string) *OpenAIModel {
	return &OpenAIModel{
		client: openai.NewClient(apiKey),
	}
}

func (m *OpenAIModel) Chat(ctx context.Context, input string) (string, error) {
	resp, err := m.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "你是智能客服"},
			{Role: "user", Content: input},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
