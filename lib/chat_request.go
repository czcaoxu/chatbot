package lib

type ChatRequest struct {
	UserID    string  `json:"user_id"`
	SessionID string  `json:"session_id"`
	Text      string  `json:"text"`
	Model     AIModel `json:"model"`
}
