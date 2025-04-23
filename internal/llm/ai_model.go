package llm

import (
	"chatbot/internal/domain"
	"context"
)

// AIModel 定义 AI 统一接口
type AIModel interface {
	ModelType() domain.AIModelType
	Chat(ctx context.Context, request *ChatRequest) (string, error)
}
