package repository

import (
	"app/internal/models"
	"errors"

	"gorm.io/gorm"
)

// ALL THE REF / DATA OBJECTS ARE TRANSFERRED TO THE REPOSITORY TO HANDLE

type NewRepositoryInterface interface {
	// GetNewsByCategory(category string) models.News
	GetNewsByCategory(category string) ([]models.News, error)
	AddNews(news models.News) error
	UpdateNews(category string, id int, news models.News) error
	DeleteNews(category string, id int) error
}

type newRepository struct {
	db *gorm.DB
}

func NewNewRepository(db *gorm.DB) *newRepository {
	return &newRepository{db: db}
}

// Read - GET
func (n *newRepository) GetNewsByCategory(category string) ([]models.News, error) {

	var News []models.News
	err := n.db.Where("category = ?", category).Find(&News).Error
	if err != nil {
		return nil, err
	}
	var newsList []models.News
	for _, news := range News {
		if news.Category == category {
			newsList = append(newsList, news)
		}
	}

	if len(newsList) == 0 {
		return nil, errors.New("no news found for the given category")
	}

	return newsList, nil
}

// Create - POST
func (n *newRepository) AddNews(news models.News) error {

	err := n.db.Create(&news).Error
	if err != nil {
		return err
	}

	return nil
}

// Update - PUT
func (n *newRepository) UpdateNews(category string, id int, newstoUpdate models.News) error {

	// Update the news
	err := n.db.Model(&newstoUpdate).Updates(newstoUpdate).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete - DELETE
func (n *newRepository) DeleteNews(category string, id int) error {

	var news models.News
	err := n.db.Where("category = ? AND id = ?", category, id).Delete(&news).Error
	if err != nil {
		return err
	}
	return nil
}
