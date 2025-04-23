package domain

type ChatInput struct {
	UserID    string      `json:"user_id"`
	SessionID string      `json:"session_id"`
	Text      string      `json:"text"`
	ModelType AIModelType `json:"model_type"`
}
