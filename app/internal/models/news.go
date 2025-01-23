package models

// GO: CamelCase, file -> snake_case
type News struct {
	ID       int    `json:"id"`
	Link     string `json:"link"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// CamelCase is used for `go`
type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type NewsCategory struct {
	NewsID     int `json:"news_id"`
	CategoryID int `json:"category_id"`
}
