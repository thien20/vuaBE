package infra

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(conStr string) *gorm.DB {

	// This is for local machine
	// "db": "root:admin123@tcp(db:3306)/vnexpress"
	// This is for docker
	// "db": "root:admin123@tcp(locahost:3306)/vnexpress"

	log.Println("Connecting to:", conStr)
	gormDB, err := gorm.Open(mysql.Open(conStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	// Take 1 value from `news` table
	// var news models.News
	// gormDB.First(&news)
	// log.Println("First news: ", news)

	log.Println("Database connected")
	// The returned `db` will include all tables within a DB
	return gormDB
}
