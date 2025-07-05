package controller

import (
	"net/http"
	"strconv"

	"synapse/internal/model"
	"synapse/internal/service"
	"synapse/internal/utils"
	"synapse/pkg/notifier"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	channelService *service.ChannelService
}

func NewChannelController(channelService *service.ChannelService) *ChannelController {
	return &ChannelController{channelService: channelService}
}

type CreateChannelRequest struct {
	Name        string                 `json:"name" binding:"required,min=1,max=255"`
	Type        string                 `json:"type" binding:"required"`
	Credentials map[string]interface{} `json:"credentials" binding:"required"`
}

// CreateChannel 创建通道
// @Summary 创建通知通道
// @Description 创建新的通知通道
// @Tags 通道
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body CreateChannelRequest true "通道信息"
// @Success 201 {object} model.Channel
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /channels [post]
func (c *ChannelController) CreateChannel(ctx *gin.Context) {
	userID, exists := ctx.Get("userId")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	var req CreateChannelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	channel := &model.Channel{
		UserID:      userID.(uint64),
		Name:        req.Name,
		Type:        req.Type,
		Credentials: model.JSON(req.Credentials),
	}

	if err := c.channelService.CreateChannel(channel); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "创建通道失败", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, channel)
}

// GetChannels 获取用户的所有通道
// @Summary 获取用户通道列表
// @Description 获取当前用户的所有通知通道
// @Tags 通道
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.Channel
// @Failure 401 {object} utils.ErrorResponse
// @Router /channels [get]
func (c *ChannelController) GetChannels(ctx *gin.Context) {
	userID, exists := ctx.Get("userId")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	channels, err := c.channelService.GetChannelsByUserID(userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "获取通道列表失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, channels)
}

// GetChannel 获取单个通道
// @Summary 获取通道详情
// @Description 根据ID获取通道详情
// @Tags 通道
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "通道ID"
// @Success 200 {object} model.Channel
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /channels/{id} [get]
func (c *ChannelController) GetChannel(ctx *gin.Context) {
	userID, exists := ctx.Get("userId")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	channel, err := c.channelService.GetChannelByID(id, userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "通道不存在", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

type UpdateChannelRequest struct {
	Name        string                 `json:"name" binding:"required,min=1,max=255"`
	Type        string                 `json:"type" binding:"required"`
	Credentials map[string]interface{} `json:"credentials" binding:"required"`
}

// UpdateChannel 更新通道
// @Summary 更新通道
// @Description 更新指定通道的信息
// @Tags 通道
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "通道ID"
// @Param data body UpdateChannelRequest true "通道信息"
// @Success 200 {object} model.Channel
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /channels/{id} [put]
func (c *ChannelController) UpdateChannel(ctx *gin.Context) {
	userID, exists := ctx.Get("userId")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	var req UpdateChannelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	channel := &model.Channel{
		ID:          id,
		UserID:      userID.(uint64),
		Name:        req.Name,
		Type:        req.Type,
		Credentials: model.JSON(req.Credentials),
	}

	if err := c.channelService.UpdateChannel(channel, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "更新通道失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

// DeleteChannel 删除通道
// @Summary 删除通道
// @Description 删除指定的通道
// @Tags 通道
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "通道ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /channels/{id} [delete]
func (c *ChannelController) DeleteChannel(ctx *gin.Context) {
	userID, exists := ctx.Get("userId")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	if err := c.channelService.DeleteChannel(id, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "删除通道失败", err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}

// TestTelegramChannel 测试 Telegram 通道
func TestTelegramChannel(ctx *gin.Context) {
	type Req struct {
		BotToken string `json:"botToken" binding:"required"`
		ChatID   string `json:"chatId" binding:"required"`
		Proxy    string `json:"proxy"`
		Content  string `json:"content" binding:"required"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, 400, "参数错误", err.Error())
		return
	}
	err := notifier.SendTelegramMessage(model.TelegramConfig{
		BotToken: req.BotToken,
		ChatID:   req.ChatID,
		Proxy:    req.Proxy,
	}, req.Content)
	if err != nil {
		utils.ErrorResponse(ctx, 500, "发送失败", err.Error())
		return
	}
	ctx.JSON(200, gin.H{"message": "发送成功"})
}

// TestEmailChannel 测试 Email 通道
func TestEmailChannel(ctx *gin.Context) {
	type Req struct {
		SMTPHost     string `json:"smtpHost"`
		SMTPPort     int    `json:"smtpPort"`
		SMTPUsername string `json:"smtpUsername"`
		SMTPPassword string `json:"smtpPassword"`
		Sender       string `json:"sender"`
		To           string `json:"to"`
		Proxy        string `json:"proxy"`
		Title        string `json:"title"`
		Content      string `json:"content"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, 400, "参数错误", err.Error())
		return
	}
	err := notifier.SendEmail(notifier.EmailConfig{
		Host:     req.SMTPHost,
		Port:     req.SMTPPort,
		Username: req.SMTPUsername,
		Password: req.SMTPPassword,
		From:     req.Sender,
		To:       req.To,
		Proxy:    req.Proxy,
	}, req.Title, req.Content)
	if err != nil {
		utils.ErrorResponse(ctx, 500, "发送失败", err.Error())
		return
	}
	ctx.JSON(200, gin.H{"message": "发送成功"})
}

// TestWebhookChannel 测试 Webhook 通道
func TestWebhookChannel(ctx *gin.Context) {
	type Req struct {
		URL     string            `json:"url"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Proxy   string            `json:"proxy"`
		Content string            `json:"content"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, 400, "参数错误", err.Error())
		return
	}
	resp, err := notifier.SendWebhook(notifier.WebhookConfig{
		URL:     req.URL,
		Method:  req.Method,
		Headers: req.Headers,
		Proxy:   req.Proxy,
	}, []byte(req.Content))
	if err != nil {
		utils.ErrorResponse(ctx, 500, "发送失败", err.Error())
		return
	}
	ctx.JSON(200, gin.H{"message": "发送成功", "response": resp})
}
