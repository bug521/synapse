package service

import (
	"errors"
	"synapse/internal/model"
	"synapse/internal/repository"

	"gorm.io/gorm"
)

type RoutingService struct {
	routingRepo *repository.RoutingRepository
	topicRepo   *repository.TopicRepository
	channelRepo *repository.ChannelRepository
}

func NewRoutingService(db *gorm.DB) *RoutingService {
	return &RoutingService{
		routingRepo: repository.NewRoutingRepository(db),
		topicRepo:   repository.NewTopicRepository(db),
		channelRepo: repository.NewChannelRepository(db),
	}
}

// CreateRouting 创建路由
func (s *RoutingService) CreateRouting(routing *model.Routing, userID uint64) error {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(routing.TopicID)
	if err != nil {
		return errors.New("主题不存在")
	}
	if topic.UserID != userID {
		return errors.New("无权访问此主题")
	}

	// 验证通道所有权
	channel, err := s.channelRepo.FindByID(routing.ChannelID)
	if err != nil {
		return errors.New("通道不存在")
	}
	if channel.UserID != userID {
		return errors.New("无权访问此通道")
	}

	// 检查路由是否已存在
	existingRouting, err := s.routingRepo.FindByTopicAndChannel(routing.TopicID, routing.ChannelID)
	if err == nil && existingRouting != nil {
		return errors.New("路由已存在")
	}

	return s.routingRepo.Create(routing)
}

// GetRoutingsByTopicID 获取主题的所有路由
func (s *RoutingService) GetRoutingsByTopicID(topicID uint64, userID uint64) ([]model.Routing, error) {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(topicID)
	if err != nil {
		return nil, errors.New("主题不存在")
	}
	if topic.UserID != userID {
		return nil, errors.New("无权访问此主题")
	}

	return s.routingRepo.FindByTopicID(topicID)
}

// GetRoutingsByChannelID 获取通道的所有路由
func (s *RoutingService) GetRoutingsByChannelID(channelID uint64, userID uint64) ([]model.Routing, error) {
	// 验证通道所有权
	channel, err := s.channelRepo.FindByID(channelID)
	if err != nil {
		return nil, errors.New("通道不存在")
	}
	if channel.UserID != userID {
		return nil, errors.New("无权访问此通道")
	}

	return s.routingRepo.FindByChannelID(channelID)
}

// UpdateRouting 更新路由
func (s *RoutingService) UpdateRouting(routing *model.Routing, userID uint64) error {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(routing.TopicID)
	if err != nil {
		return errors.New("主题不存在")
	}
	if topic.UserID != userID {
		return errors.New("无权访问此主题")
	}

	// 验证通道所有权
	channel, err := s.channelRepo.FindByID(routing.ChannelID)
	if err != nil {
		return errors.New("通道不存在")
	}
	if channel.UserID != userID {
		return errors.New("无权访问此通道")
	}

	// 检查路由是否存在
	existingRouting, err := s.routingRepo.FindByTopicAndChannel(routing.TopicID, routing.ChannelID)
	if err != nil {
		return errors.New("路由不存在")
	}

	// 保持创建时间不变
	routing.CreatedAt = existingRouting.CreatedAt

	return s.routingRepo.Update(routing)
}

// DeleteRouting 删除路由
func (s *RoutingService) DeleteRouting(topicID, channelID uint64, userID uint64) error {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(topicID)
	if err != nil {
		return errors.New("主题不存在")
	}
	if topic.UserID != userID {
		return errors.New("无权访问此主题")
	}

	// 验证通道所有权
	channel, err := s.channelRepo.FindByID(channelID)
	if err != nil {
		return errors.New("通道不存在")
	}
	if channel.UserID != userID {
		return errors.New("无权访问此通道")
	}

	return s.routingRepo.Delete(topicID, channelID)
}
