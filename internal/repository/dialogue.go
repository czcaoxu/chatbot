package repository

import (
	"chatbot/internal/domain"
	"chatbot/internal/repository/dao"
	"context"
	"fmt"
	"time"
)

// DialogueRepository
type DialogueRepository interface {
	// QueryDialogue 查询历史对话
	QueryDialogue(ctx context.Context, req *domain.ChatInput) ([]*domain.Dialogue, error)

	// CreateDialogue 创建新的对话
	CreateDialogue(ctx context.Context, req *domain.Dialogue) error
}

// dialogueRepository 对话仓储实现
type dialogueRepository struct {
	dao dao.DialogueDAO
}

func (d *dialogueRepository) QueryDialogue(ctx context.Context, req *domain.ChatInput) ([]*domain.Dialogue, error) {
	res, err := d.dao.QueryDialogue(ctx, req.UserID, req.SessionID, req.ModelType.String())
	if err != nil {
		return nil, fmt.Errorf("查询Dialogue失败: use_id: %v, session_id: %v, model:%v", req.UserID, req.SessionID, req.ModelType)
	}
	dias := make([]*domain.Dialogue, len(res))
	for i := range res {
		dias[i] = d.toDomain(res[i])
	}

	return dias, nil
}

func (d *dialogueRepository) CreateDialogue(ctx context.Context, dia *domain.Dialogue) error {
	err := d.dao.CreateDialogue(ctx, d.toEntity(dia))
	return err
}

func (d *dialogueRepository) toEntity(dia *domain.Dialogue) *dao.Dialogue {
	return &dao.Dialogue{
		ID:        dia.ID,
		UserID:    dia.UserID,
		SessionID: dia.SessionID,
		ModelType: dia.ModelType,
		Request:   dia.Request,
		Response:  dia.Response,
		CreatedAt: time.Time{},
	}
}

func (d *dialogueRepository) toDomain(dia *dao.Dialogue) *domain.Dialogue {
	return &domain.Dialogue{
		ID:        dia.ID,
		UserID:    dia.UserID,
		SessionID: dia.SessionID,
		ModelType: dia.ModelType,
		Request:   dia.Request,
		Response:  dia.Response,
	}
}

func NewDialogueRepository(d dao.DialogueDAO) DialogueRepository {
	return &dialogueRepository{
		dao: d,
	}
}
