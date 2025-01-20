package models

type News struct {
	ID      int    `json:"id"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
