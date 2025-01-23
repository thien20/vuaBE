package migration

import (
	"app/internal/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.News{})
	if err != nil {
		return err
	}

	return nil
}
