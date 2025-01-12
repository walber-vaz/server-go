package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"server-go/internal/config"
	"strings"
)

var (
	log *zap.Logger
)

func InitLogger(cfg *config.App) {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs(cfg.LogOutput)},
		Level:       zap.NewAtomicLevelAt(getLevelLogs(cfg.LogLevel)),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	err := log.Sync()
	if err != nil {
		return
	}
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	e := log.Sync()
	if e != nil {
		return
	}
}

func getOutputLogs(output string) string {
	output = strings.ToLower(strings.TrimSpace(output))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs(level string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
