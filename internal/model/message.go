package model

import (
	"time"

	"gorm.io/gorm"
)

// Message 消息模型
type Message struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;comment:消息ID" json:"id"`
	TopicID   uint64         `gorm:"not null;index;comment:来源主题ID" json:"topicId"`
	Content   JSON           `gorm:"type:json;not null;comment:原始消息内容" json:"content"`
	Status    string         `gorm:"type:varchar(50);default:'pending';index;comment:处理状态" json:"status"`
	CreatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
