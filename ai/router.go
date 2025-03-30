package ai

import "errors"

// ModelRouter 负责动态选择模型
type ModelRouter struct {
	models map[string]AIModel
}

func NewModelRouter() *ModelRouter {
	return &ModelRouter{
		models: map[string]AIModel{
			"openai": NewOpenAIModel("your-open-ai-key"),
			"qwen":   NewQwenModel("your-qwen-api-key"),
			"llama":  NewLlamaModel("your-llama-url"),
		},
	}
}

func (r *ModelRouter) GetModel(name string) (AIModel, error) {
	if model, exists := r.models[name]; exists {
		return model, nil
	}
	return nil, errors.New("模型不存在")
}
