package ai

import (
	"chatbot/config"
	"errors"
)

// ModelRouter 负责动态选择模型
type ModelRouter struct {
	models map[string]AIModel
}

func NewModelRouter(config *config.AIModelConfig) *ModelRouter {
	return &ModelRouter{
		models: map[string]AIModel{
			"openai": NewOpenAIModel(config.OpenAIKey),
			"qwen":   NewQwenModel(config.QwenKey),
			"llama":  NewLlamaModel(config.LlamaHost, config.LlamaPort),
		},
	}
}

func (r *ModelRouter) GetModel(name string) (AIModel, error) {
	if model, exists := r.models[name]; exists {
		return model, nil
	}
	return nil, errors.New("模型不存在")
}
