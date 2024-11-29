package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/nifle3/monkeyMorning/cmd/monkeyMorning/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("failed load .env file")
		os.Exit(1)
	}
	app.Run()
}
