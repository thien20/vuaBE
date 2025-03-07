package main

import (
	"app/config"
	"app/handler"
	"app/internal/cache"
	"app/internal/infra"
	"app/internal/kafka"
	"app/migration"
	"app/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.ReadConfigAndArg()
	log.Println(cfg.DB)
	database := infra.InitDB(cfg.DB)
	err := migration.Migration(database)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Redis initialization
	redisClient := config.NewRedis(cfg.Redis)
	redisCache := cache.NewCacheFromClient(redisClient)

	// News handler initialization
	newRepository := repository.NewNewRepository(database)
	newHandler := handler.NewNewsHandler(newRepository, redisCache)

	// Job handler initialization
	producer := kafka.NewKafkaProducer("scrape-news", 0) // default topic and partition - we can custom it later
	jobRepository := repository.NewJobRepository(database, producer)
	jobHandler := handler.NewJobHandler(jobRepository, redisCache)

	router := gin.Default()

	// API for news
	newsRoutes := router.Group("/news")
	{
		newsRoutes.GET("/:category", newHandler.GetNewsByCategory)
		newsRoutes.POST("/:category", newHandler.AddNews)
		newsRoutes.PUT("/:category/:id", newHandler.UpdateNews)
		newsRoutes.DELETE("/:category/:id", newHandler.DeleteNews)
	}

	// API for jobs
	jobRoutes := router.Group("/jobs")
	{
		jobRoutes.POST("/fetch/:topic", jobHandler.FetchJobs)
		jobRoutes.GET("/ping/:id", jobHandler.CheckStatus)
		jobRoutes.GET("/result/:category", jobHandler.GetResult)
	}

	// router.Run("8080:8080")
	router.Run("0.0.0.0:8080")
}
