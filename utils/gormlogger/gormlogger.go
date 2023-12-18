package gormlogger

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

type Config struct {
	LogLevel logger.LogLevel
}

type Logger struct {
	Config
}

func New(config Config) logger.Interface {
	return &Logger{
		Config: config,
	}
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level

	return l
}

// Info 替换gorm的默认日志，使用go-zero
func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Info {
		now := time.Now().Format(time.RFC3339Nano)
		callers := strings.Split(utils.FileWithLineNum(), "/")
		caller := "/" + strings.Join(callers[len(callers)-3:], "/")
		logx.WithContext(ctx).
			WithCallerSkip(4).
			Infow(
				fmt.Sprintf(s, i...),
				logx.Field("@timestamp", now),
				logx.Field("level", "info"),
				logx.Field("caller", caller),
			)
	}
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Warn {
		now := time.Now().Format(time.RFC3339Nano)
		callers := strings.Split(utils.FileWithLineNum(), "/")
		caller := "/" + strings.Join(callers[len(callers)-3:], "/")
		logx.WithContext(ctx).
			WithCallerSkip(4).
			Infow(
				fmt.Sprintf(s, i...),
				logx.Field("@timestamp", now),
				logx.Field("level", "warn"),
				logx.Field("caller", caller),
			)
	}
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Error {
		now := time.Now().Format(time.RFC3339Nano)
		callers := strings.Split(utils.FileWithLineNum(), "/")
		caller := "/" + strings.Join(callers[len(callers)-3:], "/")
		logx.WithContext(ctx).
			WithCallerSkip(4).
			Errorw(
				fmt.Sprintf(s, i...),
				logx.Field("@timestamp", now),
				logx.Field("level", "error"),
				logx.Field("caller", caller),
			)
	}
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	now := time.Now().Format(time.RFC3339Nano)
	callers := strings.Split(utils.FileWithLineNum(), "/")
	caller := "/" + strings.Join(callers[len(callers)-3:], "/")
	duration := timex.ReprOfDuration(time.Since(begin))

	if err != nil && l.LogLevel >= logger.Error {
		sql, rows := fc()
		logx.WithContext(ctx).
			WithCallerSkip(4).
			Errorw(
				"gorm happened error",
				logx.Field("@timestamp", now),
				logx.Field("level", "error"),
				logx.Field("caller", caller),
				logx.Field("rows", rows),
				logx.Field("duration", duration),
				logx.Field("sql", sql),
				logx.Field("error", fmt.Sprintf("%+v", err)),
			)
	} else {
		sql, rows := fc()
		logx.WithContext(ctx).
			WithCallerSkip(4).
			Infow(
				"gorm info",
				logx.Field("@timestamp", now),
				logx.Field("level", "error"),
				logx.Field("caller", caller),
				logx.Field("rows", rows),
				logx.Field("duration", duration),
				logx.Field("sql", sql),
			)
	}
}
