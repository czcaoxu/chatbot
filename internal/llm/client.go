package llm

import (
	"chatbot/internal/domain"
	"context"
	"errors"
	"fmt"
)

type ChatRequest struct {
	ModelType domain.AIModelType
	Dialogues *domain.Dialogues
}

type Client interface {
	Chat(ctx context.Context, request *ChatRequest) (string, error)
}

type clientImp struct {
	client map[domain.AIModelType]Client
}

func NewClientImp(clients ...AIModel) Client {
	client := make(map[domain.AIModelType]Client, len(clients))
	for _, c := range clients {
		client[c.ModelType()] = c
	}

	return &clientImp{
		client: client,
	}
}

func (c *clientImp) Chat(ctx context.Context, request *ChatRequest) (string, error) {
	if request == nil {
		return "", errors.New("nil ChatRequest")
	}

	llmClient, ok := c.client[request.ModelType]
	if !ok {
		return "", fmt.Errorf("unsupported model type %v", request.ModelType)
	}

	return llmClient.Chat(ctx, request)
}
