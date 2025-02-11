package handler

import (
	"app/internal/cache"
	"app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobRepository repository.JobRepositoryInterface
	cache         *cache.Cache
}

func NewJobHandler(jobRepository repository.JobRepositoryInterface,
	cache *cache.Cache) *JobHandler {
	return &JobHandler{
		jobRepository: jobRepository,
		cache:         cache,
	}
}

func (h *JobHandler) FetchJobs(c *gin.Context) {

	// Fetch jobs
	topic := c.Param("topic")
	if topic == "" || topic != "scrape" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Topic is required"})
		return
	}

	topic = "scrape-news" // Consumer topic for this api

	action := "scrape"
	_, err := h.jobRepository.FetchJobs(topic, action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jobs fetched successfully"})
	return
}

func (h *JobHandler) CheckStatus(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

	// Convert id to int
	id, _ := strconv.Atoi(idParam)

	// Check status
	// status, err := h.cache.Get(id)
	status, err := h.jobRepository.CheckStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get status: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
	return
}

func (h *JobHandler) GetResult(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is needed"})
		return
	}
	result, err := h.jobRepository.GetResult(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get result: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
