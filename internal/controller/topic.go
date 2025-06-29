package controller

import (
	"net/http"
	"strconv"

	"synapse/internal/model"
	"synapse/internal/service"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
)

type TopicController struct {
	topicService *service.TopicService
}

func NewTopicController(topicService *service.TopicService) *TopicController {
	return &TopicController{topicService: topicService}
}

type CreateTopicRequest struct {
	Name            string `json:"name" binding:"required,min=1,max=255"`
	SendingStrategy string `json:"sending_strategy" binding:"required"`
	ExecutionMode   string `json:"execution_mode" binding:"required"`
	Description     string `json:"description"`
}

// CreateTopic 创建主题
// @Summary 创建主题
// @Description 创建新的消息主题
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body CreateTopicRequest true "主题信息"
// @Success 201 {object} model.Topic
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /topics [post]
func (c *TopicController) CreateTopic(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	var req CreateTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	topic := &model.Topic{
		UserID:          userID.(uint64),
		Name:            req.Name,
		SendingStrategy: req.SendingStrategy,
		ExecutionMode:   req.ExecutionMode,
		Description:     req.Description,
	}

	if err := c.topicService.CreateTopic(topic); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "创建主题失败", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, topic)
}

// GetTopics 获取用户的所有主题
// @Summary 获取主题列表
// @Description 获取当前用户的所有主题
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.Topic
// @Failure 401 {object} utils.ErrorResponse
// @Router /topics [get]
func (c *TopicController) GetTopics(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	topics, err := c.topicService.GetTopicsByUserID(userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "获取主题列表失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, topics)
}

// GetTopic 获取单个主题
// @Summary 获取主题详情
// @Description 根据ID获取主题详情
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "主题ID"
// @Success 200 {object} model.Topic
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /topics/{id} [get]
func (c *TopicController) GetTopic(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	topic, err := c.topicService.GetTopicByID(id, userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "主题不存在", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, topic)
}

type UpdateTopicRequest struct {
	Name            string `json:"name" binding:"required,min=1,max=255"`
	SendingStrategy string `json:"sending_strategy" binding:"required"`
	ExecutionMode   string `json:"execution_mode" binding:"required"`
	Description     string `json:"description"`
}

// UpdateTopic 更新主题
// @Summary 更新主题
// @Description 更新指定主题的信息
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "主题ID"
// @Param data body UpdateTopicRequest true "主题信息"
// @Success 200 {object} model.Topic
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /topics/{id} [put]
func (c *TopicController) UpdateTopic(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	var req UpdateTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	topic := &model.Topic{
		ID:              id,
		UserID:          userID.(uint64),
		Name:            req.Name,
		SendingStrategy: req.SendingStrategy,
		ExecutionMode:   req.ExecutionMode,
		Description:     req.Description,
	}

	if err := c.topicService.UpdateTopic(topic, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "更新主题失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, topic)
}

// DeleteTopic 删除主题
// @Summary 删除主题
// @Description 删除指定的主题
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "主题ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /topics/{id} [delete]
func (c *TopicController) DeleteTopic(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	if err := c.topicService.DeleteTopic(id, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "删除主题失败", err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RegenerateWebhookKey 重新生成Webhook Key
// @Summary 重新生成Webhook Key
// @Description 为指定主题重新生成Webhook Key
// @Tags 主题
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "主题ID"
// @Success 200 {object} model.Topic
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /topics/{id}/regenerate-key [post]
func (c *TopicController) RegenerateWebhookKey(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	topic, err := c.topicService.RegenerateWebhookKey(id, userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "重新生成Webhook Key失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, topic)
}
