package handler

import (
	"chatbot/internal/domain"
	"errors"
	"github.com/jinzhu/copier"
)

type ChatRequest struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	Text      string `json:"text"`
	ModelType string `json:"model_type"`
}

func (c *ChatRequest) ToDomainEntity() (*domain.ChatInput, error) {
	if c == nil {
		return nil, errors.New("nil ChatRequest receiver")
	}

	var input domain.ChatInput
	if err := copier.Copy(&input, c); err != nil {
		return nil, err
	}
	input.ModelType = domain.AIModelType(c.ModelType)
	return &input, nil
}
