package main

import (
	// need for autload .env file
	_ "github.com/joho/godotenv/autoload"

	"github.com/nifle3/monkeyMorning/cmd/monkeyAfternoon/internal/app"
)

func main() {
	app.Run()
}
