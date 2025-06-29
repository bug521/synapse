package model

import (
	"time"

	"gorm.io/gorm"
)

// MessageDeliveryLog 投递日志模型
type MessageDeliveryLog struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;comment:日志ID" json:"id"`
	MessageID uint64         `gorm:"not null;index;comment:消息ID" json:"messageId"`
	ChannelID uint64         `gorm:"not null;index;comment:目标通道ID" json:"channelId"`
	Status    string         `gorm:"type:varchar(50);not null;comment:投递状态" json:"status"`
	Response  string         `gorm:"type:text;comment:API响应" json:"response"`
	CreatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
