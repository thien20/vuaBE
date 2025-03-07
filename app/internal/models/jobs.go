package models

// GO: CamelCase, file -> snake_case

// CamelCase is used for `go`

type Jobs struct {
	ID     uint   `gorm:"primaryKey"`
	Status string `gorm:"type:varchar(50)"`
}
