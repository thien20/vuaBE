package migration

import (
	"app/internal/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.News{},
		&models.Jobs{},
		&models.ScrapedResult{})
	if err != nil {
		return err
	}

	return nil
}
