package repository

import (
	"app/models"
	"database/sql"
	"errors"
)

// ALL THE REF / DATA OBJECTS ARE TRANSFERRED TO THE REPOSITORY TO HANDLE

type NewRepositoryInterface interface {
	// GetNewsByCategory(category string) models.News
	GetNewsByCategory(category string) ([]models.News, error)
	AddNews(category string, news models.News) error
	UpdateNews(category string, id int, news models.News) error
	DeleteNews(category string, id int) error
}

type newRepository struct {
	db *sql.DB
}

func NewNewRepository(db *sql.DB) *newRepository {
	return &newRepository{db: db}
}

// Read - GET
func (n *newRepository) GetNewsByCategory(category string) ([]models.News, error) {
	query := "SELECT id, link, title, content FROM `" + category + "`"
	rows, err := n.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID, &news.Link, &news.Title, &news.Content); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

// Create - POST
func (n *newRepository) AddNews(category string, news models.News) error {
	query := "INSERT INTO `" + category + "` (id, link, title, content) VALUES (?, ?, ?, ?)"
	_, err := n.db.Exec(query, news.ID, news.Link, news.Title, news.Content)
	if err != nil {
		return err
	}

	return nil
}

// Update - PUT
func (n *newRepository) UpdateNews(category string, id int, news models.News) error {
	query := "UPDATE `" + category + "` SET link = ?, title = ?, content = ? WHERE id = ?"
	result, err := n.db.Exec(query, news.Link, news.Title, news.Content, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

// Delete - DELETE
func (n *newRepository) DeleteNews(category string, id int) error {
	query := "DELETE FROM `" + category + "` WHERE id = ?"
	result, err := n.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
