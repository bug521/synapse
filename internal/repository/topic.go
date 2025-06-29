package repository

import (
	"crypto/rand"
	"encoding/hex"
	"synapse/internal/model"

	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

// Create 创建主题
func (r *TopicRepository) Create(topic *model.Topic) error {
	// 生成唯一的Webhook Key
	topic.WebhookKey = r.GenerateWebhookKey()
	return r.db.Create(topic).Error
}

// FindByID 根据ID查找主题
func (r *TopicRepository) FindByID(id uint64) (*model.Topic, error) {
	var topic model.Topic
	err := r.db.First(&topic, id).Error
	return &topic, err
}

// FindByUserID 根据用户ID查找所有主题
func (r *TopicRepository) FindByUserID(userID uint64) ([]model.Topic, error) {
	var topics []model.Topic
	err := r.db.Where("user_id = ?", userID).Find(&topics).Error
	return topics, err
}

// FindByWebhookKey 根据Webhook Key查找主题
func (r *TopicRepository) FindByWebhookKey(webhookKey string) (*model.Topic, error) {
	var topic model.Topic
	err := r.db.Where("webhook_key = ?", webhookKey).First(&topic).Error
	return &topic, err
}

// Update 更新主题
func (r *TopicRepository) Update(topic *model.Topic) error {
	return r.db.Save(topic).Error
}

// Delete 删除主题
func (r *TopicRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Topic{}, id).Error
}

// DeleteByUserID 删除用户的所有主题
func (r *TopicRepository) DeleteByUserID(userID uint64) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Topic{}).Error
}

// GenerateWebhookKey 生成唯一的Webhook Key
func (r *TopicRepository) GenerateWebhookKey() string {
	for {
		bytes := make([]byte, 16)
		rand.Read(bytes)
		webhookKey := hex.EncodeToString(bytes)

		// 检查是否已存在
		var count int64
		r.db.Model(&model.Topic{}).Where("webhook_key = ?", webhookKey).Count(&count)
		if count == 0 {
			return webhookKey
		}
	}
}
