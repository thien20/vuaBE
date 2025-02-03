package handler

import (
	"app/internal/cache"
	"app/internal/models"
	"app/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HANDLER IS USED TO HANDLE BUSINESS LOGIC

type NewsHandler struct {
	newRepository repository.NewRepositoryInterface
	cache         *cache.Cache
}

func NewNewsHandler(newRepository repository.NewRepositoryInterface,
	cache *cache.Cache) *NewsHandler {
	return &NewsHandler{
		newRepository: newRepository,
		cache:         cache,
	}
}

func (h *NewsHandler) GetNewsByCategory(c *gin.Context) {
	category := c.Param("category")
	// cacheKey := "news:" + category

	// Check if the data is in cache
	cachedValue, err := h.cache.GetCache(category)
	if err == nil {
		log.Println("Cache hit for category:", category)
		c.JSON(http.StatusOK, cachedValue)
		return
	}

	// If not in cache, fetch from database
	log.Println("Cache miss for category:", category)

	newsList, err := h.newRepository.GetNewsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news: " + err.Error()})
		return
	}

	// Convert newsList to JSON
	newsListJSON, err := json.Marshal(newsList)
	if err != nil {
		log.Println("Failed to marshal newsList for category:", category)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process news data"})
		return
	}

	// Cache the data
	err = h.cache.SetCache(category, newsListJSON, 10000000000)
	if err != nil {
		log.Println("Failed to cache data for category:", category)
	}

	c.JSON(http.StatusOK, newsList)
}

func (h *NewsHandler) AddNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	err := h.newRepository.AddNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert news: " + err.Error()})
		return
	}

	// Invalidate cache for the category
	err = h.cache.DeleteCache(news.Category)
	if err != nil {
		log.Println("Failed to invalidate cache for category:", news.Category)
	}

	c.JSON(http.StatusCreated, news)
}

func (h *NewsHandler) UpdateNews(c *gin.Context) {
	category := c.Param("category")
	idParam := c.Param("id")
	// The param take the id as string, so we need to convert it to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	var newstoUpdate models.News
	if err := c.ShouldBindJSON(&newstoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	err = h.newRepository.UpdateNews(category, id, newstoUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news: " + err.Error()})
		return
	}

	// Invalidate cache for the category
	err = h.cache.DeleteCache(category)
	if err != nil {
		log.Println("Failed to invalidate cache for category:", category)
	}

	c.JSON(http.StatusOK, gin.H{"message": "News updated successfully"})
}

func (h *NewsHandler) DeleteNews(c *gin.Context) {
	category := c.Param("category")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	err = h.newRepository.DeleteNews(category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news: " + err.Error()})
		return
	}

	// Invalidate cache for the category
	err = h.cache.DeleteCache(category)
	if err != nil {
		log.Println("Failed to invalidate cache for category:", category)
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
