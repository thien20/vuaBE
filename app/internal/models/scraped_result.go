package models

import "time"

type ScrapedResult struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	JobID     uint      `gorm:"not null;index"`
	Category  string    `gorm:"type:varchar(255);not null"`
	Link      string    `gorm:"type:text"` //;not null"`
	Title     string    `gorm:"type:text"` //;not null"`
	Content   string    `gorm:"type:text"` //;not null"`
	ScrapedAt time.Time `json:"scraped_at"`
}
