package model

import (
	"time"

	"gorm.io/gorm"
)

type Channel struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;comment:通道ID"`
	UserID      uint64         `gorm:"not null;index;comment:所属用户的ID"`
	Name        string         `gorm:"type:varchar(255);not null;comment:通道名称"`
	Type        string         `gorm:"type:varchar(50);not null;comment:通道类型"`
	Credentials JSON           `gorm:"type:json;not null;comment:凭证"` // 使用自定义JSON类型（需实现Scanner/Valuer接口）
	CreatedAt   time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Telegram 配置（结构化）
type TelegramConfig struct {
	BotToken  string `json:"bot_token"`
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

// email 配置（结构化）
type EmailConfig struct {
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password"`
	Sender       string `json:"sender"`
}
