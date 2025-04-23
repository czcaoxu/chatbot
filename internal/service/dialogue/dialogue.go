package dialogue

import (
	"chatbot/internal/domain"
	"chatbot/internal/llm"
	"chatbot/internal/repository"
	"context"
	"fmt"
)

type Service interface {
	Chat(ctx context.Context, chatInput *domain.ChatInput) (string, error)
}

type dialogueService struct {
	repo   repository.DialogueRepository
	client llm.Client
}

func NewDialogueService(repo repository.DialogueRepository, client llm.Client) Service {
	return &dialogueService{
		repo:   repo,
		client: client,
	}
}

func (d *dialogueService) Chat(ctx context.Context, input *domain.ChatInput) (string, error) {
	// 获取历史记录
	dialogues, err := d.repo.QueryDialogue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("query dialogues failed: %v", err.Error())
	}
	dialogues = append(dialogues, &domain.Dialogue{
		Request: input.Text,
	})
	chatRequest := &llm.ChatRequest{
		ModelType: input.ModelType,
		Dialogues: &domain.Dialogues{
			SystemRole:   "assistant",
			UserRole:     "user",
			DialogueList: dialogues,
		},
	}

	reply, err := d.client.Chat(ctx, chatRequest)
	if err != nil {
		return "", fmt.Errorf("chat with model failed: %v", err.Error())
	}

	if err := d.repo.CreateDialogue(ctx, &domain.Dialogue{
		UserID:    input.UserID,
		SessionID: input.SessionID,
		ModelType: input.ModelType.String(),
		Request:   input.Text,
		Response:  reply,
	}); err != nil {
		fmt.Println("CreateDialogue user_id[%v], session_id[%v], mode_type[%v] failed: %v",
			input.UserID, input.SessionID, input.ModelType.String(), err.Error())
		return reply, nil
	}

	return reply, nil
}
