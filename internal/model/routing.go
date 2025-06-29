package model

import (
	"time"

	"gorm.io/gorm"
)

// Routing 路由模型
type Routing struct {
	TopicID          uint64         `gorm:"primaryKey;comment:项目ID" json:"topicId"`
	ChannelID        uint64         `gorm:"primaryKey;comment:通道ID" json:"channelId"`
	Priority         int            `gorm:"default:0;comment:优先级" json:"priority"`
	VariableMappings JSON           `gorm:"type:json;comment:变量映射规则" json:"variableMappings"`
	MessageTemplate  string         `gorm:"type:text;comment:消息模板" json:"messageTemplate"`
	CreatedAt        time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
