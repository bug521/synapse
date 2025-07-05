package model

import (
	"time"

	"gorm.io/gorm"
)

type Channel struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;comment:通道ID" json:"id"`
	UserID      uint64         `gorm:"not null;index;comment:所属用户的ID" json:"userId"`
	Name        string         `gorm:"type:varchar(255);not null;comment:通道名称" json:"name"`
	Type        string         `gorm:"type:varchar(50);not null;comment:通道类型" json:"type"`
	Credentials JSON           `gorm:"type:json;not null;comment:凭证" json:"credentials"` // 使用自定义JSON类型（需实现Scanner/Valuer接口）
	CreatedAt   time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Telegram 配置（结构化）
type TelegramConfig struct {
	BotToken  string `json:"botToken"`
	ChatID    string `json:"chatId"`
	ParseMode string `json:"parseMode"`
	Proxy     string `json:"proxy"`
}

// email 配置（结构化）
type EmailConfig struct {
	SMTPHost     string `json:"smtpHost"`
	SMTPPort     int    `json:"smtpPort"`
	SMTPUsername string `json:"smtpUsername"`
	SMTPPassword string `json:"smtpPassword"`
	Sender       string `json:"sender"`
	To           string `json:"to"`
}
