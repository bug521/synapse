package router

import (
	"synapse/internal/controller"
	"synapse/internal/middleware"
	"synapse/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// 创建服务
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	// 初始化Gin
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", userController.GetProfile)
		protected.PUT("/profile", userController.UpdateProfile)
		protected.DELETE("/profile", userController.DeleteAccount)
	}

	return r
}
