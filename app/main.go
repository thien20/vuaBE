package main

import (
	"log"
	"app/db"
	"app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.DB.Close()

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
