package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/nifle3/monkeyMorning/internal/config"
)

func Run() {
	ctx := context.Background()

	_, err := config.New()
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("ERROR: load config %s", err.Error()))
		os.Exit(1)
	}
}
