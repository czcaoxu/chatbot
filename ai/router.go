package ai

import (
	"chatbot/lib"
	"errors"
)

// ModelRouter 负责动态选择模型
type ModelRouter struct {
	models map[lib.AIModel]AIModel
}

func NewModelRouter() *ModelRouter {
	return &ModelRouter{
		models: map[lib.AIModel]AIModel{
			lib.ModelGPT4:   NewOpenAIModel(""),
			lib.ModelQwen:   NewQwenModel(""),
			lib.ModelLLama3: NewLlamaModel(),
		},
	}
}

func (r *ModelRouter) GetModel(name lib.AIModel) (AIModel, error) {
	if model, exists := r.models[name]; exists {
		return model, nil
	}
	return nil, errors.New("模型不存在")
}
