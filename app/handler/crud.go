package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type News struct {
	ID      int    `json:"id"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsHandler struct {
	db *sql.DB
}

func NewNewsHandler(db *sql.DB) *NewsHandler {
	return &NewsHandler{db: db}
}

// Read
func (h *NewsHandler) GetNewsByCategory(c *gin.Context) {
	category := c.Param("category")

	query := "SELECT * FROM `" + category + "`"
	rows, err := h.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve news: " + err.Error()})
		return
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var news News
		if err := rows.Scan(&news.ID, &news.Link, &news.Title, &news.Content); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse news: " + err.Error()})
			return
		}
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process rows: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, newsList)
}

// Create
func (h *NewsHandler) AddNews(c *gin.Context) {
	category := c.Param("category")
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	query := "INSERT INTO `" + category + "` (id, link, title, content) VALUES (?, ?, ?, ?)"
	_, err := h.db.Exec(query, news.ID, news.Link, news.Title, news.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert news: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

// Update
func (h *NewsHandler) UpdateNews(c *gin.Context) {
	category := c.Param("category")
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	query := "UPDATE `" + category + "` SET link = ?, title = ?, content = ? WHERE id = ?"
	result, err := h.db.Exec(query, news.Link, news.Title, news.Content, news.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rows affected: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// Delete
func (h *NewsHandler) DeleteNews(c *gin.Context) {
	category := c.Param("category")
	id := c.Param("id")

	query := "DELETE FROM `" + category + "` WHERE id = ?"
	result, err := h.db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rows affected: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
