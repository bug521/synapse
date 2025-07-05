package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"sort"
	"synapse/internal/model"
	"synapse/internal/repository"
	"synapse/pkg/notifier"

	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

type MessageService struct {
	messageRepo  *repository.MessageRepository
	topicRepo    *repository.TopicRepository
	routingRepo  *repository.RoutingRepository
	channelRepo  *repository.ChannelRepository
	deliveryRepo *repository.DeliveryRepository
}

func NewMessageService(db *gorm.DB) *MessageService {
	return &MessageService{
		messageRepo:  repository.NewMessageRepository(db),
		topicRepo:    repository.NewTopicRepository(db),
		routingRepo:  repository.NewRoutingRepository(db),
		channelRepo:  repository.NewChannelRepository(db),
		deliveryRepo: repository.NewDeliveryRepository(db),
	}
}

// CreateMessage 创建消息
func (s *MessageService) CreateMessage(message *model.Message) error {
	return s.messageRepo.Create(message)
}

// GetMessageByID 根据ID获取消息
func (s *MessageService) GetMessageByID(id uint64) (*model.Message, error) {
	return s.messageRepo.FindByID(id)
}

// GetMessagesByTopicID 获取主题的消息列表
func (s *MessageService) GetMessagesByTopicID(topicID uint64, page, pageSize int) ([]model.Message, int64, error) {
	offset := (page - 1) * pageSize
	messages, err := s.messageRepo.FindByTopicID(topicID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.messageRepo.CountByTopicID(topicID)
	if err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

// ProcessMessage 处理消息
func (s *MessageService) ProcessMessage(messageID uint64) error {
	// 获取消息
	message, err := s.messageRepo.FindByID(messageID)
	if err != nil {
		return err
	}

	// 更新消息状态为处理中
	if err := s.messageRepo.UpdateStatus(messageID, "processing"); err != nil {
		return err
	}

	// 获取主题
	topic, err := s.topicRepo.FindByID(message.TopicID)
	if err != nil {
		s.messageRepo.UpdateStatus(messageID, "failed")
		return err
	}

	// 获取路由规则
	routings, err := s.routingRepo.FindByTopicID(message.TopicID)
	if err != nil {
		s.messageRepo.UpdateStatus(messageID, "failed")
		return err
	}

	if len(routings) == 0 {
		// 没有路由规则，标记为完成
		s.messageRepo.UpdateStatus(messageID, "completed")
		return nil
	}

	// 根据发送策略处理消息
	switch topic.SendingStrategy {
	case "all":
		return s.processAllStrategy(message, routings)
	case "failover":
		return s.processFailoverStrategy(message, routings)
	default:
		s.messageRepo.UpdateStatus(messageID, "failed")
		return errors.New("不支持的发送策略")
	}
}

// processAllStrategy 处理"发送给所有"策略
func (s *MessageService) processAllStrategy(message *model.Message, routings []model.Routing) error {
	successCount := 0
	totalCount := len(routings)

	for _, routing := range routings {
		if err := s.sendToChannel(message, &routing); err != nil {
			// 记录失败日志，但继续处理其他通道
			s.logDeliveryFailure(message.ID, routing.ChannelID, err.Error())
		} else {
			successCount++
			s.logDeliverySuccess(message.ID, routing.ChannelID)
		}
	}

	// 更新消息状态
	if successCount == 0 {
		s.messageRepo.UpdateStatus(message.ID, "failed")
	} else if successCount == totalCount {
		s.messageRepo.UpdateStatus(message.ID, "completed")
	} else {
		s.messageRepo.UpdateStatus(message.ID, "partial")
	}

	return nil
}

// processFailoverStrategy 处理"故障转移"策略
func (s *MessageService) processFailoverStrategy(message *model.Message, routings []model.Routing) error {
	// sort routings by priority
	sort.Slice(routings, func(i, j int) bool {
		return routings[i].Priority > routings[j].Priority
	})

	for _, routing := range routings {
		if err := s.sendToChannel(message, &routing); err != nil {
			// 记录失败日志，继续尝试下一个通道
			s.logDeliveryFailure(message.ID, routing.ChannelID, err.Error())
			continue
		} else {
			// 发送成功，记录日志并结束
			s.logDeliverySuccess(message.ID, routing.ChannelID)
			s.messageRepo.UpdateStatus(message.ID, "completed")
			return nil
		}
	}

	// 所有通道都失败了
	s.messageRepo.UpdateStatus(message.ID, "failed")
	return errors.New("所有通道发送失败")
}

// sendToChannel 发送消息到指定通道
func (s *MessageService) sendToChannel(message *model.Message, routing *model.Routing) error {
	// 获取通道信息
	channel, err := s.channelRepo.FindByID(routing.ChannelID)
	if err != nil {
		return err
	}

	// 根据通道类型发送消息
	switch channel.Type {
	case "telegram":
		return s.sendToTelegram(message, channel, routing)
	case "email":
		return s.sendToEmail(message, channel, routing)
	case "slack":
		return s.sendToSlack(message, channel, routing)
	case "webhook":
		return s.sendToWebhook(message, channel, routing)
	default:
		return errors.New("不支持的通道类型")
	}
}

// sendToTelegram 发送到Telegram
func (s *MessageService) sendToTelegram(message *model.Message, channel *model.Channel, routing *model.Routing) error {
	var cfg model.TelegramConfig
	b, _ := json.Marshal(channel.Credentials)
	if err := json.Unmarshal(b, &cfg); err != nil {
		return err
	}
	contentBytes, _ := json.Marshal(message.Content)

	variables := make(map[string]interface{})
	for name, path := range routing.VariableMappings {
		pathStr, ok := path.(string)
		if !ok {
			continue
		}
		variables[name] = gjson.GetBytes(contentBytes, pathStr).Value()
	}
	tmpl, err := template.New("message").Parse(routing.MessageTemplate)
	if err != nil {
		return err
	}
	var renderedMessage bytes.Buffer
	if err := tmpl.Execute(&renderedMessage, variables); err != nil {
		return err
	}

	return notifier.SendTelegramMessage(model.TelegramConfig{
		BotToken:  cfg.BotToken,
		ChatID:    cfg.ChatID,
		Proxy:     cfg.Proxy,
		ParseMode: cfg.ParseMode,
	}, string(renderedMessage.Bytes()))
}

// sendToEmail 发送到Email
func (s *MessageService) sendToEmail(message *model.Message, channel *model.Channel, routing *model.Routing) error {
	var cfg struct {
		SMTPHost     string `json:"smtpHost"`
		SMTPPort     int    `json:"smtpPort"`
		SMTPUsername string `json:"smtpUsername"`
		SMTPPassword string `json:"smtpPassword"`
		Sender       string `json:"sender"`
		To           string `json:"to"`
		Proxy        string `json:"proxy"`
	}
	b, _ := json.Marshal(channel.Credentials)
	if err := json.Unmarshal(b, &cfg); err != nil {
		return err
	}
	contentBytes, _ := json.Marshal(message.Content)

	messageContent := string(contentBytes)
	return notifier.SendEmail(notifier.EmailConfig{
		Host:     cfg.SMTPHost,
		Port:     cfg.SMTPPort,
		Username: cfg.SMTPUsername,
		Password: cfg.SMTPPassword,
		From:     cfg.Sender,
		To:       cfg.To,
		Proxy:    cfg.Proxy,
	}, "test", messageContent)
}

// sendToSlack 发送到Slack
func (s *MessageService) sendToSlack(message *model.Message, channel *model.Channel, routing *model.Routing) error {
	// TODO: 实现Slack发送逻辑
	// 1. 解析Slack配置
	// 2. 渲染消息模板
	// 3. 调用Slack API
	return nil
}

// sendToWebhook 发送到Webhook
func (s *MessageService) sendToWebhook(message *model.Message, channel *model.Channel, routing *model.Routing) error {
	var cfg struct {
		URL     string            `json:"url"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Proxy   string            `json:"proxy"`
	}
	b, _ := json.Marshal(channel.Credentials)
	if err := json.Unmarshal(b, &cfg); err != nil {
		return err
	}
	contentBytes, _ := json.Marshal(message.Content)

	resp, err := notifier.SendWebhook(notifier.WebhookConfig{
		URL:     cfg.URL,
		Method:  cfg.Method,
		Headers: cfg.Headers,
		Proxy:   cfg.Proxy,
	}, contentBytes)
	if err != nil {
		return err
	}
	_ = resp // 可记录响应内容
	return nil
}

// logDeliverySuccess 记录发送成功日志
func (s *MessageService) logDeliverySuccess(messageID, channelID uint64) {
	deliveryLog := &model.MessageDeliveryLog{
		MessageID: messageID,
		ChannelID: channelID,
		Status:    "success",
		Response:  "发送成功",
	}
	s.deliveryRepo.Create(deliveryLog)
}

// logDeliveryFailure 记录发送失败日志
func (s *MessageService) logDeliveryFailure(messageID, channelID uint64, errorMsg string) {
	deliveryLog := &model.MessageDeliveryLog{
		MessageID: messageID,
		ChannelID: channelID,
		Status:    "failed",
		Response:  errorMsg,
	}
	s.deliveryRepo.Create(deliveryLog)
}
