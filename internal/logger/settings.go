package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/nifle3/monkeyMorning/internal/config"
)

func Setting(cfg *config.Logger) (io.Closer, error) {
	logFile, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	level := getLevelByString(cfg.Level)
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	handler := slog.NewJSONHandler(logFile, opts)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logFile, nil
}

func getLevelByString(level string) slog.Leveler {
	level = strings.ToLower(level)

	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "error":
		return slog.LevelError
	}

	return slog.LevelDebug
}
