package logger

import (
	"log/slog"
	"os"
	"time"
)

type LoggerConfig struct {
	BaseLogLevel slog.Level
	Logger       slog.Logger
}

type Option func(*LoggerConfig)

func NewLoggerConfig(opts ...Option) *LoggerConfig {
	lc := &LoggerConfig{}
	for _, opt := range opts {
		opt(lc)
	}
	return lc
}

func withBaseLogLevel(level slog.Level) Option {
	return func(lc *LoggerConfig) {
		lc.BaseLogLevel = level
	}
}

func withSlog() Option {
	return func(lc *LoggerConfig) {
		log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		lc.Logger = *log
	}
}

func ConvertLogLevel(level string) slog.Level {
	switch level {
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}
	return slog.LevelInfo
}

func DetermineLogLevel(status int) slog.Level {
	if status == 200 {
		return slog.LevelInfo
	}
	return slog.LevelError
}

func ConvertLatency(latency time.Duration) string {
	return latency.String()
}

func InitLogger(level string) *LoggerConfig {
	lv := ConvertLogLevel(level)
	lc := NewLoggerConfig(withBaseLogLevel(lv), withSlog())
	return lc
}
