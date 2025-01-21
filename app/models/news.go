package models

type News struct {
	ID       int    `json:"id"`
	Link     string `json:"link"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type Category struct {
	Category_ID   int    `json:"category_id"`
	Category_name string `json:"category_name"`
}

type News_Category struct {
	News_ID     int `json:"news_id"`
	Category_ID int `json:"category_id"`
}
