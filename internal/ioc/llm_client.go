package ioc

import (
	"chatbot/internal/llm"
	"chatbot/internal/llm/llama"
	"chatbot/internal/llm/openai"
	"chatbot/internal/llm/qwen"
)

func InitLLMClient() llm.Client {
	return llm.NewClientImp(llama.NewLlamaModel(),
		openai.NewOpenAIModel(),
		qwen.NewQwenModel(),
	)
}
