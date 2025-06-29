package repository

import (
	"synapse/internal/model"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

// Create 创建投递日志
func (r *DeliveryRepository) Create(deliveryLog *model.MessageDeliveryLog) error {
	return r.db.Create(deliveryLog).Error
}

// FindByMessageID 根据消息ID查找投递日志
func (r *DeliveryRepository) FindByMessageID(messageID uint64) ([]model.MessageDeliveryLog, error) {
	var logs []model.MessageDeliveryLog
	err := r.db.Where("message_id = ?", messageID).Order("created_at DESC").Find(&logs).Error
	return logs, err
}

// FindByChannelID 根据通道ID查找投递日志
func (r *DeliveryRepository) FindByChannelID(channelID uint64, limit, offset int) ([]model.MessageDeliveryLog, error) {
	var logs []model.MessageDeliveryLog
	err := r.db.Where("channel_id = ?", channelID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, err
}

// CountByChannelID 统计通道的投递日志数量
func (r *DeliveryRepository) CountByChannelID(channelID uint64) (int64, error) {
	var count int64
	err := r.db.Model(&model.MessageDeliveryLog{}).Where("channel_id = ?", channelID).Count(&count).Error
	return count, err
}

// FindByStatus 根据状态查找投递日志
func (r *DeliveryRepository) FindByStatus(status string, limit, offset int) ([]model.MessageDeliveryLog, error) {
	var logs []model.MessageDeliveryLog
	err := r.db.Where("status = ?", status).Order("created_at DESC").Limit(limit).Offset(offset).Find(&logs).Error
	return logs, err
}

// DeleteByMessageID 删除消息的所有投递日志
func (r *DeliveryRepository) DeleteByMessageID(messageID uint64) error {
	return r.db.Where("message_id = ?", messageID).Delete(&model.MessageDeliveryLog{}).Error
}

// DeleteByChannelID 删除通道的所有投递日志
func (r *DeliveryRepository) DeleteByChannelID(channelID uint64) error {
	return r.db.Where("channel_id = ?", channelID).Delete(&model.MessageDeliveryLog{}).Error
}
