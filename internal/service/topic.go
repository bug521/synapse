package service

import (
	"errors"
	"synapse/internal/model"
	"synapse/internal/repository"

	"gorm.io/gorm"
)

type TopicService struct {
	topicRepo *repository.TopicRepository
}

func NewTopicService(db *gorm.DB) *TopicService {
	return &TopicService{
		topicRepo: repository.NewTopicRepository(db),
	}
}

// CreateTopic 创建主题
func (s *TopicService) CreateTopic(topic *model.Topic) error {
	// 验证发送策略
	if !s.isValidSendingStrategy(topic.SendingStrategy) {
		return errors.New("不支持的发送策略")
	}

	// 验证执行模式
	if !s.isValidExecutionMode(topic.ExecutionMode) {
		return errors.New("不支持的执行模式")
	}

	return s.topicRepo.Create(topic)
}

// GetTopicByID 根据ID获取主题
func (s *TopicService) GetTopicByID(id uint64, userID uint64) (*model.Topic, error) {
	topic, err := s.topicRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 验证主题所有权
	if topic.UserID != userID {
		return nil, errors.New("无权访问此主题")
	}

	return topic, nil
}

// GetTopicsByUserID 获取用户的所有主题
func (s *TopicService) GetTopicsByUserID(userID uint64) ([]model.Topic, error) {
	return s.topicRepo.FindByUserID(userID)
}

// GetTopicByWebhookKey 根据Webhook Key获取主题
func (s *TopicService) GetTopicByWebhookKey(webhookKey string) (*model.Topic, error) {
	return s.topicRepo.FindByWebhookKey(webhookKey)
}

// UpdateTopic 更新主题
func (s *TopicService) UpdateTopic(topic *model.Topic, userID uint64) error {
	// 验证主题所有权
	existingTopic, err := s.topicRepo.FindByID(topic.ID)
	if err != nil {
		return err
	}

	if existingTopic.UserID != userID {
		return errors.New("无权修改此主题")
	}

	// 验证发送策略
	if !s.isValidSendingStrategy(topic.SendingStrategy) {
		return errors.New("不支持的发送策略")
	}

	// 验证执行模式
	if !s.isValidExecutionMode(topic.ExecutionMode) {
		return errors.New("不支持的执行模式")
	}

	topic.UserID = userID                       // 确保用户ID不被修改
	topic.WebhookKey = existingTopic.WebhookKey // 保持Webhook Key不变
	return s.topicRepo.Update(topic)
}

// DeleteTopic 删除主题
func (s *TopicService) DeleteTopic(id uint64, userID uint64) error {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(id)
	if err != nil {
		return err
	}

	if topic.UserID != userID {
		return errors.New("无权删除此主题")
	}

	return s.topicRepo.Delete(id)
}

// RegenerateWebhookKey 重新生成Webhook Key
func (s *TopicService) RegenerateWebhookKey(id uint64, userID uint64) (*model.Topic, error) {
	// 验证主题所有权
	topic, err := s.topicRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if topic.UserID != userID {
		return nil, errors.New("无权修改此主题")
	}

	// 生成新的Webhook Key
	topic.WebhookKey = s.topicRepo.GenerateWebhookKey()

	if err := s.topicRepo.Update(topic); err != nil {
		return nil, err
	}

	return topic, nil
}

// isValidSendingStrategy 验证发送策略是否有效
func (s *TopicService) isValidSendingStrategy(strategy string) bool {
	validStrategies := []string{"all", "failover"}
	for _, s := range validStrategies {
		if s == strategy {
			return true
		}
	}
	return false
}

// isValidExecutionMode 验证执行模式是否有效
func (s *TopicService) isValidExecutionMode(mode string) bool {
	validModes := []string{"async", "sync"}
	for _, m := range validModes {
		if m == mode {
			return true
		}
	}
	return false
}
