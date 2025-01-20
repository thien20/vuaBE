package main

import (
	"app/db"
	"app/handler"
	"app/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB initialization
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	defer database.Close()

	newRepository := repository.NewNewRepository(database)
	handler := handler.NewNewsHandler(newRepository)

	router := gin.Default()

	newsRoutes := router.Group("/news")
	{
		newsRoutes.GET("/:category", handler.GetNewsByCategory)
		newsRoutes.POST("/:category", handler.AddNews)
		newsRoutes.PUT("/:category/:id", handler.UpdateNews)
		newsRoutes.DELETE("/:category/:id", handler.DeleteNews)
	}

	log.Println("Server started on http://localhost:8080")
	router.Run("localhost:8080")
}
