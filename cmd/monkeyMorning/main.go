package main

import (
	// need for autoload .env file
	_ "github.com/joho/godotenv/autoload"

	"github.com/nifle3/monkeyMorning/cmd/monkeyMorning/internal/app"
)

func main() {
	app.Run()
}
