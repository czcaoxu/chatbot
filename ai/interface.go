package ai

import "context"

// AIModel 定义 AI 统一接口
type AIModel interface {
	Chat(ctx context.Context, input string) (string, error)
}
