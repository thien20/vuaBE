package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error

	// dsn := "root:admin123@tcp(localhost:3306)/news"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	// Check if the database is connected
	// rows, err := DB.Query("SELECT id, title FROM the_gioi LIMIT 1")
	// if err != nil {
	// 	return err
	// }
	// defer rows.Close()

	// var id int
	// var title string

	// if rows.Next() {
	// 	err = rows.Scan(&id, &title)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Printf("Fetched article: ID=%d, Title=%s", id, title)
	// } else {
	// 	log.Println("No articles found")
	// }

	log.Println("Database connected!")
	return nil
}
