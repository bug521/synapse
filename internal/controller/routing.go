package controller

import (
	"net/http"
	"strconv"

	"synapse/internal/model"
	"synapse/internal/service"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
)

type RoutingController struct {
	routingService *service.RoutingService
}

func NewRoutingController(routingService *service.RoutingService) *RoutingController {
	return &RoutingController{routingService: routingService}
}

type CreateRoutingRequest struct {
	TopicID          uint64                 `json:"topic_id" binding:"required"`
	ChannelID        uint64                 `json:"channel_id" binding:"required"`
	Priority         int                    `json:"priority"`
	VariableMappings map[string]interface{} `json:"variable_mappings"`
	MessageTemplate  string                 `json:"message_template"`
}

// CreateRouting 创建路由
// @Summary 创建路由规则
// @Description 创建新的消息路由规则
// @Tags 路由
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body CreateRoutingRequest true "路由信息"
// @Success 201 {object} model.Routing
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /routings [post]
func (c *RoutingController) CreateRouting(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	var req CreateRoutingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	routing := &model.Routing{
		TopicID:          req.TopicID,
		ChannelID:        req.ChannelID,
		Priority:         req.Priority,
		VariableMappings: model.JSON(req.VariableMappings),
		MessageTemplate:  req.MessageTemplate,
	}

	if err := c.routingService.CreateRouting(routing, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "创建路由失败", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, routing)
}

// GetRoutingsByTopic 获取主题的所有路由
// @Summary 获取主题路由列表
// @Description 获取指定主题的所有路由规则
// @Tags 路由
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "主题ID"
// @Success 200 {array} model.Routing
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /topics/{id}/routings [get]
func (c *RoutingController) GetRoutingsByTopic(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	topicIDStr := ctx.Param("id")
	topicID, err := strconv.ParseUint(topicIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	routings, err := c.routingService.GetRoutingsByTopicID(topicID, userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "获取路由列表失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, routings)
}

// GetRoutingsByChannel 获取通道的所有路由
// @Summary 获取通道路由列表
// @Description 获取指定通道的所有路由规则
// @Tags 路由
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "通道ID"
// @Success 200 {array} model.Routing
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /channels/{id}/routings [get]
func (c *RoutingController) GetRoutingsByChannel(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	channelIDStr := ctx.Param("id")
	channelID, err := strconv.ParseUint(channelIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	routings, err := c.routingService.GetRoutingsByChannelID(channelID, userID.(uint64))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "获取路由列表失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, routings)
}

type UpdateRoutingRequest struct {
	Priority         int                    `json:"priority"`
	VariableMappings map[string]interface{} `json:"variable_mappings"`
	MessageTemplate  string                 `json:"message_template"`
}

// UpdateRouting 更新路由
// @Summary 更新路由规则
// @Description 更新指定的路由规则
// @Tags 路由
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param channel_id path int true "通道ID"
// @Param data body UpdateRoutingRequest true "路由信息"
// @Success 200 {object} model.Routing
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /routings/{topic_id}/{channel_id} [put]
func (c *RoutingController) UpdateRouting(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	topicIDStr := ctx.Param("topic_id")
	topicID, err := strconv.ParseUint(topicIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	channelIDStr := ctx.Param("channel_id")
	channelID, err := strconv.ParseUint(channelIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	var req UpdateRoutingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	routing := &model.Routing{
		TopicID:          topicID,
		ChannelID:        channelID,
		Priority:         req.Priority,
		VariableMappings: model.JSON(req.VariableMappings),
		MessageTemplate:  req.MessageTemplate,
	}

	if err := c.routingService.UpdateRouting(routing, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "更新路由失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, routing)
}

// DeleteRouting 删除路由
// @Summary 删除路由规则
// @Description 删除指定的路由规则
// @Tags 路由
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param channel_id path int true "通道ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /routings/{topic_id}/{channel_id} [delete]
func (c *RoutingController) DeleteRouting(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	topicIDStr := ctx.Param("topic_id")
	topicID, err := strconv.ParseUint(topicIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的主题ID", err.Error())
		return
	}

	channelIDStr := ctx.Param("channel_id")
	channelID, err := strconv.ParseUint(channelIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "无效的通道ID", err.Error())
		return
	}

	if err := c.routingService.DeleteRouting(topicID, channelID, userID.(uint64)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "删除路由失败", err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
