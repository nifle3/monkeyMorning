package app

import (
	"log/slog"
	"os"

	"github.com/nifle3/monkeyMorning/internal/config"
)

func Run() {
	_, err := config.New()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// logger
	// s3
	// domain
	// bot
}
