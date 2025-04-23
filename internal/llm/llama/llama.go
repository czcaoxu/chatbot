package llama

import (
	"bytes"
	"chatbot/internal/domain"
	"chatbot/internal/llm"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const LlamaAPI = "/v1/chat/completions"

// LlamaModel 适配本地 LLaMA
type LlamaModel struct {
	apiURL string
}

func NewLlamaModel() *LlamaModel {
	url := os.Getenv("OLLAMA_URL") + LlamaAPI
	fmt.Println("llama url: ", url)
	return &LlamaModel{apiURL: url}
}

func (m *LlamaModel) ModelType() domain.AIModelType {
	return domain.ModelLLama3
}

func (m *LlamaModel) Chat(ctx context.Context, request *llm.ChatRequest) (string, error) {
	dialogues := request.Dialogues.ToDialogueModel()
	message := make([]map[string]string, len(dialogues)+1)
	message[0] = map[string]string{
		"role":    "system",
		"content": "You are a helpful assistant.",
	}
	copy(message[1:], dialogues)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model":    "llama3",
		"messages": message,
	})
	resp, err := http.Post(m.apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("读取响应错误: " + err.Error())
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		fmt.Println("JSON 解析错误:", err)
		fmt.Println("Body:", body)
		return "", errors.New("JSON 解析错误: " + err.Error())
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", errors.New("LLaMA API 解析失败")
}
