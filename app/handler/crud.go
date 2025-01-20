package handler

import (
	"app/models"
	"app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HANDLER IS USED TO HANDLE BUSINESS LOGIC

type NewsHandler struct {
	newRepository repository.NewRepositoryInterface
}

func NewNewsHandler(newRepository repository.NewRepositoryInterface) *NewsHandler {
	return &NewsHandler{newRepository: newRepository}
}

func (h *NewsHandler) GetNewsByCategory(c *gin.Context) {
	category := c.Param("category")
	newsList, err := h.newRepository.GetNewsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, newsList)
}

func (h *NewsHandler) AddNews(c *gin.Context) {
	category := c.Param("category")
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	err := h.newRepository.AddNews(category, news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert news: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

func (h *NewsHandler) UpdateNews(c *gin.Context) {
	category := c.Param("category")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + idParam})
		return
	}

	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	err = h.newRepository.UpdateNews(category, id, news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *NewsHandler) DeleteNews(c *gin.Context) {
	category := c.Param("category")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + idParam})
		return
	}

	err = h.newRepository.DeleteNews(category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
