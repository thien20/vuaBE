package repository

import (
	"app/internal/kafka"
	"app/internal/models"

	"gorm.io/gorm"
)

type JobRepositoryInterface interface {
	FetchJobs(topic string, action string) ([]models.Jobs, error)
	CheckStatus(id int) ([]models.Jobs, error)
	GetResult(category string) ([]models.ScrapedResult, error)
}

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (j *JobRepository) FetchJobs(topic string, action string) ([]models.Jobs, error) {
	// Initialize the producer
	key := "Scrape-Key"
	id := 1
	producer, err := kafka.NewProducer(topic)
	if err != nil {
		return nil, err
	}

	// Produce the message
	err = producer.ProduceMessage(id, key, action)
	if err != nil {
		return nil, err
	}

	err = j.db.Create(&models.Jobs{
		Status: "pending"}).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}

// Check Status from `jobs` table
func (j *JobRepository) CheckStatus(id int) ([]models.Jobs, error) {

	var status []models.Jobs
	// SELECT `status` FROM `jobs` WHERE `id` = input_id
	err := j.db.Select("status").Where("id = ?", id).Find(&status).Error
	// err := j.db.Where("id = ?", id).Find(&status).Error
	if err != nil {
		return nil, err
	}

	return status, nil
}

// Get Result from `scraped_results` table
func (j *JobRepository) GetResult(category string) ([]models.ScrapedResult, error) {
	var result []models.ScrapedResult
	err := j.db.Where("category = ?", category).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
