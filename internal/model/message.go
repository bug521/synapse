package model

import (
	"time"

	"gorm.io/gorm"
)

// Message 消息模型
type Message struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;comment:消息ID"`
	TopicID   uint64         `gorm:"not null;index;comment:来源主题ID"`
	Content   JSON           `gorm:"type:json;not null;comment:原始消息内容"`
	Status    string         `gorm:"type:varchar(50);default:'pending';index;comment:处理状态"`
	CreatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
