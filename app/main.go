package main

import (
	"app/config"
	"app/handler"
	"app/internal/cache"
	"app/internal/infra"
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

	// Repository and handler initialization
	newRepository := repository.NewNewRepository(database)
	handler := handler.NewNewsHandler(newRepository, redisCache)

	router := gin.Default()

	newsRoutes := router.Group("/news")
	{
		newsRoutes.GET("/:category", handler.GetNewsByCategory)
		newsRoutes.POST("/:category", handler.AddNews)
		newsRoutes.PUT("/:category/:id", handler.UpdateNews)
		newsRoutes.DELETE("/:category/:id", handler.DeleteNews)
	}

	// router.Run("8080:8080")
	router.Run("0.0.0.0:8080")
}
