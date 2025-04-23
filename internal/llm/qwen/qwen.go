package qwen

import (
	"chatbot/internal/domain"
	"chatbot/internal/llm"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type QwenModel struct {
	apiKey string
}

func NewQwenModel() *QwenModel {
	return &QwenModel{}
}

func (m *QwenModel) ModelType() domain.AIModelType {
	return domain.ModelQwen
}

func (m *QwenModel) Chat(ctx context.Context, request *llm.ChatRequest) (string, error) {
	input := ""
	reqBody := `{"prompt": "` + input + `", "model": "qwen-max"}`
	req, _ := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/qwen/chat", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "Bearer "+m.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["output"].(string), nil
}
