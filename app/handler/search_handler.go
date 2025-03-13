package handler

import (
	"app/repository"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	repo repository.SearchRepositoryInterface
}

func NewSearchHandler(repo repository.SearchRepositoryInterface) *SearchHandler {
	return &SearchHandler{repo: repo}
}

func (h *SearchHandler) SearchSimple(c *gin.Context) {
	wordToSearch := c.Param("simple")
	results, err := h.repo.SearchSimple(wordToSearch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"results": results})
}
