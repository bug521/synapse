package repository

import (
	"synapse/internal/model"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create 创建消息
func (r *MessageRepository) Create(message *model.Message) error {
	return r.db.Create(message).Error
}

// FindByID 根据ID查找消息
func (r *MessageRepository) FindByID(id uint64) (*model.Message, error) {
	var message model.Message
	err := r.db.First(&message, id).Error
	return &message, err
}

// FindByTopicID 根据主题ID查找消息
func (r *MessageRepository) FindByTopicID(topicID uint64, limit, offset int) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Where("topic_id = ?", topicID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&messages).Error
	return messages, err
}

// CountByTopicID 统计主题的消息数量
func (r *MessageRepository) CountByTopicID(topicID uint64) (int64, error) {
	var count int64
	err := r.db.Model(&model.Message{}).Where("topic_id = ?", topicID).Count(&count).Error
	return count, err
}

// UpdateStatus 更新消息状态
func (r *MessageRepository) UpdateStatus(id uint64, status string) error {
	return r.db.Model(&model.Message{}).Where("id = ?", id).Update("status", status).Error
}

// Delete 删除消息
func (r *MessageRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Message{}, id).Error
}

// DeleteByTopicID 删除主题的所有消息
func (r *MessageRepository) DeleteByTopicID(topicID uint64) error {
	return r.db.Where("topic_id = ?", topicID).Delete(&model.Message{}).Error
}
