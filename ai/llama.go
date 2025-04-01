package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const LlamaAPI = "/v1/chat/completions"

// LlamaModel 适配本地 LLaMA
type LlamaModel struct {
	apiURL string
}

func NewLlamaModel(host, port string) *LlamaModel {
	return &LlamaModel{apiURL: "http://" + host + ":" + port + LlamaAPI}
}

func (m *LlamaModel) Chat(ctx context.Context, input string) (string, error) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "llama3",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a helpful assistant.",
			},
			{
				"role":    "user",
				"content": input,
			},
		},
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
		return "", errors.New("JSON 解析错误: " + err.Error())
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", errors.New("LLaMA API 解析失败")
}

type ChatResponse struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	Created           int64  `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
