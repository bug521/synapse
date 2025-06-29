package model

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement;comment:项目ID"`
	UserID          uint64         `gorm:"not null;index;comment:所属用户的ID"`
	Name            string         `gorm:"type:varchar(255);not null;comment:项目名称"`
	WebhookKey      string         `gorm:"type:varchar(36);not null;uniqueIndex;comment:Webhook Key"`
	SendingStrategy string         `gorm:"type:varchar(50);default:'all';comment:发送策略"`
	ExecutionMode   string         `gorm:"type:varchar(50);default:'async';comment:执行模式"`
	Description     string         `gorm:"type:text;comment:项目描述"`
	CreatedAt       time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间"`
	UpdatedAt       time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
