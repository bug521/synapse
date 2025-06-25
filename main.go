package main

import (
	"fmt"
	"log"
	"net/http"
	"synapse/internal/config"
	"synapse/internal/model"
	"synapse/internal/router"
	"synapse/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. 加载配置
	config.InitConfig("config/config.yaml")
	cfg := config.GlobalConfig

	// 2. 初始化日志
	logger.InitLogger()
	defer logger.Logger.Sync()
	zap.L().Info("配置加载完成", zap.Any("config", cfg))

	// 3. 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	zap.L().Info("数据库连接成功")

	// 4. 自动迁移数据模型
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	zap.L().Info("数据库迁移完成")

	// 5. 设置Gin模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 6. 初始化路由
	r := router.SetupRouter(db)

	// 7. 启动服务器
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	zap.L().Info("服务器启动中", zap.String("address", serverAddr))

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("服务器启动失败", zap.Error(err))
	}
}
