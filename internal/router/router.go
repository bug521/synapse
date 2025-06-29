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
	channelService := service.NewChannelService(db)
	topicService := service.NewTopicService(db)
	routingService := service.NewRoutingService(db)
	messageService := service.NewMessageService(db)

	// 创建控制器
	userController := controller.NewUserController(userService)
	channelController := controller.NewChannelController(channelService)
	topicController := controller.NewTopicController(topicService)
	routingController := controller.NewRoutingController(routingService)
	webhookController := controller.NewWebhookController(topicService, messageService)

	// 初始化Gin
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.RecoveryMiddleware())

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	// Webhook路由（无需认证）
	webhook := r.Group("/webhook")
	{
		webhook.POST("/:webhook_key", webhookController.ReceiveWebhook)
		webhook.GET("/:webhook_key/info", webhookController.GetWebhookInfo)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户相关
		protected.GET("/profile", userController.GetProfile)
		protected.PUT("/profile", userController.UpdateProfile)
		protected.DELETE("/profile", userController.DeleteAccount)

		// 通道相关
		channels := protected.Group("/channels")
		{
			channels.POST("", channelController.CreateChannel)
			channels.GET("", channelController.GetChannels)
			channels.GET("/:id", channelController.GetChannel)
			channels.PUT("/:id", channelController.UpdateChannel)
			channels.DELETE("/:id", channelController.DeleteChannel)
			// 通道路由相关
			channels.GET("/:id/routings", routingController.GetRoutingsByChannel)
		}

		// 主题相关
		topics := protected.Group("/topics")
		{
			topics.POST("", topicController.CreateTopic)
			topics.GET("", topicController.GetTopics)
			topics.GET("/:id", topicController.GetTopic)
			topics.PUT("/:id", topicController.UpdateTopic)
			topics.DELETE("/:id", topicController.DeleteTopic)
			topics.POST("/:id/regenerate-key", topicController.RegenerateWebhookKey)
			// 主题路由相关
			topics.GET("/:id/routings", routingController.GetRoutingsByTopic)
		}

		// 路由相关
		routings := protected.Group("/routings")
		{
			routings.POST("", routingController.CreateRouting)
			routings.PUT("/:topic_id/:channel_id", routingController.UpdateRouting)
			routings.DELETE("/:topic_id/:channel_id", routingController.DeleteRouting)
		}
	}

	return r
}
