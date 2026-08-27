package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level, format string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	parsedLevel, err := zap.ParseAtomicLevel(strings.ToLower(level))
	if err != nil {
		return nil, fmt.Errorf("invalid log level %q: %w", level, err)
	}
	cfg.Level = parsedLevel
	cfg.DisableCaller = true
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	switch strings.ToLower(format) {
	case "json":
		cfg.Encoding = "json"
	case "text", "console":
		cfg.Encoding = "console"
	default:
		return nil, fmt.Errorf("invalid log format %q", format)
	}

	return cfg.Build()
}
