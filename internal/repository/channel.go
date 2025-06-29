package repository

import (
	"synapse/internal/model"

	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{db: db}
}

// Create 创建通道
func (r *ChannelRepository) Create(channel *model.Channel) error {
	return r.db.Create(channel).Error
}

// FindByID 根据ID查找通道
func (r *ChannelRepository) FindByID(id uint64) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.First(&channel, id).Error
	return &channel, err
}

// FindByUserID 根据用户ID查找所有通道
func (r *ChannelRepository) FindByUserID(userID uint64) ([]model.Channel, error) {
	var channels []model.Channel
	err := r.db.Where("user_id = ?", userID).Find(&channels).Error
	return channels, err
}

// Update 更新通道
func (r *ChannelRepository) Update(channel *model.Channel) error {
	return r.db.Save(channel).Error
}

// Delete 删除通道
func (r *ChannelRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Channel{}, id).Error
}

// DeleteByUserID 删除用户的所有通道
func (r *ChannelRepository) DeleteByUserID(userID uint64) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Channel{}).Error
}
