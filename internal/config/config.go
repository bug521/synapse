package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	Charset      string
	MaxOpenConns int
	MaxIdleConns int
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type LogConfig struct {
	Level string
	Path  string
}

var GlobalConfig Config

func InitConfig(configPath string) {
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// 创建日志目录
	logDir := filepath.Dir(GlobalConfig.Log.Path)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Printf("failed to create log directory: %v", err)
		}
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			log.Printf("Error reloading config: %v", err)
		}
	})
}
