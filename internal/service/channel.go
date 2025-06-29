package service

import (
	"encoding/json"
	"errors"
	"synapse/internal/model"
	"synapse/internal/repository"

	"gorm.io/gorm"
)

type ChannelService struct {
	channelRepo *repository.ChannelRepository
}

func NewChannelService(db *gorm.DB) *ChannelService {
	return &ChannelService{
		channelRepo: repository.NewChannelRepository(db),
	}
}

// CreateChannel 创建通道
func (s *ChannelService) CreateChannel(channel *model.Channel) error {
	// 验证通道类型
	if !s.isValidChannelType(channel.Type) {
		return errors.New("不支持的通道类型")
	}

	// 验证凭证格式
	if err := s.validateCredentials(channel.Type, channel.Credentials); err != nil {
		return err
	}

	return s.channelRepo.Create(channel)
}

// GetChannelByID 根据ID获取通道
func (s *ChannelService) GetChannelByID(id uint64, userID uint64) (*model.Channel, error) {
	channel, err := s.channelRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 验证通道所有权
	if channel.UserID != userID {
		return nil, errors.New("无权访问此通道")
	}

	return channel, nil
}

// GetChannelsByUserID 获取用户的所有通道
func (s *ChannelService) GetChannelsByUserID(userID uint64) ([]model.Channel, error) {
	return s.channelRepo.FindByUserID(userID)
}

// UpdateChannel 更新通道
func (s *ChannelService) UpdateChannel(channel *model.Channel, userID uint64) error {
	// 验证通道所有权
	existingChannel, err := s.channelRepo.FindByID(channel.ID)
	if err != nil {
		return err
	}

	if existingChannel.UserID != userID {
		return errors.New("无权修改此通道")
	}

	// 验证通道类型
	if !s.isValidChannelType(channel.Type) {
		return errors.New("不支持的通道类型")
	}

	// 验证凭证格式
	if err := s.validateCredentials(channel.Type, channel.Credentials); err != nil {
		return err
	}

	channel.UserID = userID // 确保用户ID不被修改
	return s.channelRepo.Update(channel)
}

// DeleteChannel 删除通道
func (s *ChannelService) DeleteChannel(id uint64, userID uint64) error {
	// 验证通道所有权
	channel, err := s.channelRepo.FindByID(id)
	if err != nil {
		return err
	}

	if channel.UserID != userID {
		return errors.New("无权删除此通道")
	}

	return s.channelRepo.Delete(id)
}

// isValidChannelType 验证通道类型是否有效
func (s *ChannelService) isValidChannelType(channelType string) bool {
	validTypes := []string{"telegram", "email", "slack", "webhook"}
	for _, t := range validTypes {
		if t == channelType {
			return true
		}
	}
	return false
}

// validateCredentials 验证凭证格式
func (s *ChannelService) validateCredentials(channelType string, credentials model.JSON) error {
	switch channelType {
	case "telegram":
		var config model.TelegramConfig
		credentialsBytes, err := json.Marshal(credentials)
		if err != nil {
			return errors.New("Telegram配置格式错误")
		}
		if err := json.Unmarshal(credentialsBytes, &config); err != nil {
			return errors.New("Telegram配置格式错误")
		}
		if config.BotToken == "" || config.ChatID == "" {
			return errors.New("Telegram配置不完整")
		}
	case "email":
		var config model.EmailConfig
		credentialsBytes, err := json.Marshal(credentials)
		if err != nil {
			return errors.New("Email配置格式错误")
		}
		if err := json.Unmarshal(credentialsBytes, &config); err != nil {
			return errors.New("Email配置格式错误")
		}
		if config.SMTPHost == "" || config.SMTPUsername == "" || config.SMTPPassword == "" {
			return errors.New("Email配置不完整")
		}
	}
	return nil
}
