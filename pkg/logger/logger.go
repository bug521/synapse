package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(logLevel string, logPath string) {

	// 设置日志级别
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zapcore.InfoLevel
	}

	// 创建日志核心
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 日志输出：控制台和文件
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	// 文件日志轮转
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, // MB
		MaxBackups: 7,
		MaxAge:     30, // days
		Compress:   true,
	})
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, level)

	// 组合核心
	core := zapcore.NewTee(consoleCore, fileCore)

	// 创建日志器
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 替换全局日志器
	zap.ReplaceGlobals(Logger)
}
