package main

import (
	"fmt"
	"log"

	"github.com/MiroslavArbor/student-organisations-api/internal/config"
	"github.com/MiroslavArbor/student-organisations-api/internal/db"
	"github.com/MiroslavArbor/student-organisations-api/internal/logger"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system environment variables")
	}
	config, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	var logger = logger.NewLogger(config.LOG_LEVEL)

	database, err := db.InitDB(config)

	if err != nil {
		logger.Error("Failed to initialize database:", err)
		return
	}

	if config.ENV == "development" {
		if err := db.SeedDatabase(database, config); err != nil {
			logger.Error("Failed to seed database:", err)
		}
	}

	fmt.Println("Database initialized successfully:", database)
}
