package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DialogueDAO interface {
	QueryDialogue(ctx context.Context, userId, sessionId, modelType string) ([]*Dialogue, error)
	CreateDialogue(ctx context.Context, dialogue *Dialogue) error
}

// Dialogue 结构体对应数据库表
type Dialogue struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"size:255;not null;index:idx_user_session_model"`
	SessionID string    `gorm:"size:255;not null;index:idx_user_session_model"` // 联合索引
	ModelType string    `gorm:"size:50;not null;index:idx_user_session_model"`  // 联合索引
	Request   string    `gorm:"type:text;not null"`
	Response  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type dialogueDAO struct {
	db *gorm.DB
}

func (d *dialogueDAO) QueryDialogue(ctx context.Context, userId, sessionId, modelType string) ([]*Dialogue, error) {
	var dialogues []*Dialogue
	err := d.db.WithContext(ctx).Where("user_id = ? AND session_id = ? AND model_type = ?", userId, sessionId, modelType).
		Order("created_at asc").
		Find(&dialogues).Error
	if err != nil {
		return nil, fmt.Errorf("查询dialogues失败: %v", err.Error())
	}

	return dialogues, nil
}

func (d *dialogueDAO) CreateDialogue(ctx context.Context, dialogue *Dialogue) error {
	if err := d.db.WithContext(ctx).Create(dialogue).Error; err != nil {
		return fmt.Errorf("插入dialogue失败，error：%v", err.Error())
	}

	return nil
}

func NewDialogueDAO(db *gorm.DB) DialogueDAO {
	return &dialogueDAO{
		db: db,
	}
}
