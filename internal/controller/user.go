package controller

import (
	"net/http"

	"synapse/internal/model"
	"synapse/internal/service"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Email    string `json:"email" binding:"required,email"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "注册信息"
// @Success 201 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := c.userService.Register(user); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "注册失败", err.Error())
		return
	}

	// 返回用户信息时隐藏密码
	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录获取JWT令牌
// @Tags 用户
// @Accept json
// @Produce json
// @Param data body LoginRequest true "登录凭证"
// @Success 200 {object} object{token=string}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	user, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "登录失败", "用户名或密码错误")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "生成令牌失败", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// GetProfile 获取用户信息
// @Summary 获取用户信息
// @Description 获取当前登录用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.User
// @Failure 401 {object} utils.ErrorResponse
// @Router /profile [get]
func (c *UserController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	user, err := c.userService.GetUserProfile(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "获取用户信息失败", err.Error())
		return
	}

	// 返回用户信息时隐藏密码
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

// UpdateProfile 更新用户信息
// @Summary 更新用户信息
// @Description 更新当前登录用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.User true "用户信息"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /profile [put]
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "参数错误", err.Error())
		return
	}

	// 确保更新的是自己的账户
	user.ID = userID.(uint)

	if err := c.userService.UpdateUser(&user); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "更新失败", err.Error())
		return
	}

	// 返回用户信息时隐藏密码
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

// DeleteAccount 删除用户账户
// @Summary 删除用户账户
// @Description 删除当前登录用户账户
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 204
// @Failure 401 {object} utils.ErrorResponse
// @Router /profile [delete]
func (c *UserController) DeleteAccount(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "用户信息不存在")
		return
	}

	if err := c.userService.DeleteUser(userID.(uint)); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "删除失败", err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
