package lib

import "time"

// Message 结构体对应数据库表
type Message struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"size:255;not null"`
	Message   string    `gorm:"type:text;not null"`
	Response  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
