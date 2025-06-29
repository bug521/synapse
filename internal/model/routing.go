package model

import (
	"time"

	"gorm.io/gorm"
)

// Routing 路由模型
type Routing struct {
	TopicID          uint64         `gorm:"primaryKey;comment:项目ID"`
	ChannelID        uint64         `gorm:"primaryKey;comment:通道ID"`
	Priority         int            `gorm:"default:0;comment:优先级"`
	VariableMappings JSON           `gorm:"type:json;comment:变量映射规则"`
	MessageTemplate  string         `gorm:"type:text;comment:消息模板"`
	CreatedAt        time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:接收时间"`
	UpdatedAt        time.Time      `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3);index;comment:更新时间"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
