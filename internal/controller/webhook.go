package controller

import (
	"net/http"

	"synapse/internal/model"
	"synapse/internal/service"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	topicService   *service.TopicService
	messageService *service.MessageService
}

func NewWebhookController(topicService *service.TopicService, messageService *service.MessageService) *WebhookController {
	return &WebhookController{
		topicService:   topicService,
		messageService: messageService,
	}
}

// ReceiveWebhook 接收Webhook消息
// @Summary 接收Webhook消息
// @Description 接收来自外部服务的Webhook消息
// @Tags Webhook
// @Accept json
// @Produce json
// @Param webhook_key path string true "Webhook Key"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /webhook/{webhook_key} [post]
func (c *WebhookController) ReceiveWebhook(ctx *gin.Context) {
	webhookKey := ctx.Param("webhook_key")
	if webhookKey == "" {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的Webhook Key", "Webhook Key不能为空")
		return
	}

	// 查找对应的主题
	topic, err := c.topicService.GetTopicByWebhookKey(webhookKey)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "主题不存在", "无效的Webhook Key")
		return
	}

	// 读取请求体
	var payload map[string]interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的JSON格式", err.Error())
		return
	}

	// 创建消息记录
	message := &model.Message{
		TopicID: topic.ID,
		Content: model.JSON(payload),
		Status:  "pending",
	}

	if err := c.messageService.CreateMessage(message); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "保存消息失败", err.Error())
		return
	}

	// 根据执行模式处理消息
	if topic.ExecutionMode == "sync" {
		// 同步处理
		if err := c.messageService.ProcessMessage(message.ID); err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "处理消息失败", err.Error())
			return
		}
	} else {
		// 异步处理 - 这里可以发送到消息队列
		// TODO: 实现异步消息处理
		go func() {
			c.messageService.ProcessMessage(message.ID)
		}()
	}

	// 返回成功响应
	response := map[string]interface{}{
		"message_id": message.ID,
		"status":     "received",
		"topic":      topic.Name,
	}

	ctx.JSON(http.StatusOK, response)
}

// GetWebhookInfo 获取Webhook信息
// @Summary 获取Webhook信息
// @Description 获取指定Webhook Key对应的主题信息
// @Tags Webhook
// @Accept json
// @Produce json
// @Param webhook_key path string true "Webhook Key"
// @Success 200 {object} model.Topic
// @Failure 404 {object} utils.ErrorResponse
// @Router /webhook/{webhook_key}/info [get]
func (c *WebhookController) GetWebhookInfo(ctx *gin.Context) {
	webhookKey := ctx.Param("webhook_key")
	if webhookKey == "" {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的Webhook Key", "Webhook Key不能为空")
		return
	}

	// 查找对应的主题
	topic, err := c.topicService.GetTopicByWebhookKey(webhookKey)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "主题不存在", "无效的Webhook Key")
		return
	}

	ctx.JSON(http.StatusOK, topic)
}
