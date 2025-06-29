package repository

import (
	"synapse/internal/model"

	"gorm.io/gorm"
)

type RoutingRepository struct {
	db *gorm.DB
}

func NewRoutingRepository(db *gorm.DB) *RoutingRepository {
	return &RoutingRepository{db: db}
}

// Create 创建路由
func (r *RoutingRepository) Create(routing *model.Routing) error {
	return r.db.Create(routing).Error
}

// FindByTopicID 根据主题ID查找所有路由
func (r *RoutingRepository) FindByTopicID(topicID uint64) ([]model.Routing, error) {
	var routings []model.Routing
	err := r.db.Where("topic_id = ?", topicID).Order("priority DESC").Find(&routings).Error
	return routings, err
}

// FindByChannelID 根据通道ID查找所有路由
func (r *RoutingRepository) FindByChannelID(channelID uint64) ([]model.Routing, error) {
	var routings []model.Routing
	err := r.db.Where("channel_id = ?", channelID).Find(&routings).Error
	return routings, err
}

// FindByTopicAndChannel 根据主题ID和通道ID查找路由
func (r *RoutingRepository) FindByTopicAndChannel(topicID, channelID uint64) (*model.Routing, error) {
	var routing model.Routing
	err := r.db.Where("topic_id = ? AND channel_id = ?", topicID, channelID).First(&routing).Error
	return &routing, err
}

// Update 更新路由
func (r *RoutingRepository) Update(routing *model.Routing) error {
	return r.db.Save(routing).Error
}

// Delete 删除路由
func (r *RoutingRepository) Delete(topicID, channelID uint64) error {
	return r.db.Where("topic_id = ? AND channel_id = ?", topicID, channelID).Delete(&model.Routing{}).Error
}

// DeleteByTopicID 删除主题的所有路由
func (r *RoutingRepository) DeleteByTopicID(topicID uint64) error {
	return r.db.Where("topic_id = ?", topicID).Delete(&model.Routing{}).Error
}

// DeleteByChannelID 删除通道的所有路由
func (r *RoutingRepository) DeleteByChannelID(channelID uint64) error {
	return r.db.Where("channel_id = ?", channelID).Delete(&model.Routing{}).Error
}
