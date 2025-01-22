package db

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*gorm.DB, error) {

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

	// db, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	err = db.AutoMigrate(&models.News{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected and auto migration completed!")
	// The returned `db` will include all tables within a DB
	return db, nil
}
