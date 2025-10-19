package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

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

	// Get underlying SQL DB for connection management
	sqlDB, err := database.DB()
	if err != nil {
		logger.Error("Failed to get underlying SQL DB:", err)
		return
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			logger.Error("Failed to close database connection:", err)
		} else {
			logger.Info("Database connection closed successfully")
		}
	}()

	if config.ENV == "development" {
		if err := db.SeedDatabase(database, config); err != nil {
			logger.Error("Failed to seed database:", err)
		}
	}

	fmt.Println("Database initialized successfully:", database)

	router := gin.New()
	if config.ENV == "development" {
		router.Use(gin.Logger()) // more verbose logging
	}
	router.Use(gin.Recovery())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is working!",
		})
	})

	router.Run(":3001")

}
