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
	redisClient, err := config.NewRedis()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	redisCache := cache.NewCacheFromClient(redisClient)

	err = redisCache.SetCache("the_gioi", "The gioi", 0)
	if err != nil {
		log.Fatalf("Failed to set value to Redis: %v", err)
	}
	cachedValue, err := redisCache.GetCache("the_gioi")
	if err != nil {
		log.Fatalf("Failed to get value from Redis: %v", err)
	}
	log.Println("Cached value: ", cachedValue)

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
