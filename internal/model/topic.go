package model

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement;comment:项目ID" json:"id"`
	UserID          uint64         `gorm:"not null;index;comment:所属用户的ID" json:"userId"`
	Name            string         `gorm:"type:varchar(255);not null;comment:项目名称" json:"name"`
	WebhookKey      string         `gorm:"type:varchar(36);not null;uniqueIndex;comment:Webhook Key" json:"webhookKey"`
	SendingStrategy string         `gorm:"type:varchar(50);default:'all';comment:发送策略" json:"sendingStrategy"`
	ExecutionMode   string         `gorm:"type:varchar(50);default:'async';comment:执行模式" json:"executionMode"`
	Description     string         `gorm:"type:text;comment:项目描述" json:"description"`
	CreatedAt       time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
