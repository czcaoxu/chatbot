package openai

import (
	"chatbot/internal/domain"
	"chatbot/internal/llm"
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
	client *openai.Client
}

func NewOpenAIModel() *OpenAIModel {
	return &OpenAIModel{
		client: openai.NewClient(""),
	}
}

func (m *OpenAIModel) ModelType() domain.AIModelType {
	return domain.ModelGPT4
}

func (m *OpenAIModel) Chat(ctx context.Context, request *llm.ChatRequest) (string, error) {
	resp, err := m.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "你是智能客服"},
			{Role: "user", Content: ""},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
