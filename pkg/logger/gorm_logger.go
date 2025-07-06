package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type ZapGormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

// NewZapGormLogger 创建新的GORM-Zap日志适配器
func NewZapGormLogger(zapLogger *zap.Logger, logLevel string) *ZapGormLogger {
	logLevel = strings.ToLower(logLevel)
	var logLevelInt logger.LogLevel
	switch logLevel {
	case "info":
		logLevelInt = logger.Info
	case "warn":
		logLevelInt = logger.Warn
	case "error":
		logLevelInt = logger.Error
	default:
		logLevelInt = logger.Silent
	}

	return &ZapGormLogger{
		ZapLogger:                 zapLogger,
		LogLevel:                  logLevelInt,            // 默认日志级别
		SlowThreshold:             200 * time.Millisecond, // 慢查询阈值
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true, // 默认忽略RecordNotFound错误
	}
}

// LogMode 设置日志级别
func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 打印信息
func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.ZapLogger.Info(fmt.Sprintf(msg, data...))
	}
}

// Warn 打印警告
func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.ZapLogger.Warn(fmt.Sprintf(msg, data...))
	}
}

// Error 打印错误
func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.ZapLogger.Error(fmt.Sprintf(msg, data...))
	}
}

// Trace 打印SQL跟踪
func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		l.ZapLogger.Error("gorm trace",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.ZapLogger.Warn("gorm slow sql",
			zap.Duration("elapsed", elapsed),
			zap.Duration("threshold", l.SlowThreshold),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.ZapLogger.Info("gorm sql",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	}
}

// WithContext 实现Context方法
func (l *ZapGormLogger) WithContext(ctx context.Context) logger.Interface {
	return l
}
