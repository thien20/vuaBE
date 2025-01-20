package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {

	// dsn := "root:admin123@tcp(localhost:3306)/news"
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Check if the database is connected
	// rows, err := db.Query("SELECT id, title FROM the_gioi LIMIT 1")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var id int
	// var title string

	// if rows.Next() {
	// 	err = rows.Scan(&id, &title)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	log.Printf("Fetched article: ID=%d, Title=%s", id, title)
	// } else {
	// 	log.Println("No articles found")
	// }

	log.Println("Database connected!")
	return db, nil
}
