package handler

import (
	"app/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

type News struct {
	ID      int    `json:"id"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Read
func GetNewsByCategory(c *gin.Context) {
	category := c.Param("category")
	rows, err := db.DB.Query("SELECT id, link, title, content FROM " + category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var news News
		if err := rows.Scan(&news.ID, &news.Link, &news.Title, &news.Content); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		newsList = append(newsList, news)
	}
	c.JSON(http.StatusOK, newsList)
}

// Create
func AddNews(c *gin.Context) {
	category := c.Param("category")
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("INSERT INTO "+category+" (link, title, content) VALUES (?, ?, ?)",
		news.Link, news.Title, news.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "News added successfully!"})
}

// Update
func UpdateNews(c *gin.Context) {
	category := c.Param("category")
	id := c.Param("id")
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE "+category+" SET link = ?, title = ?, content = ? WHERE id = ?",
		news.Link, news.Title, news.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News updated successfully!"})
}

// Delete
func DeleteNews(c *gin.Context) {
	category := c.Param("category")
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM "+category+" WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully!"})
}
