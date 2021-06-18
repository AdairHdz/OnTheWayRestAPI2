package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func init() {
	databaseUsername := os.Getenv("USERNAME")
	databasePassword := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databaseAddress := os.Getenv("DB_ADDRESS")
	databasePort := os.Getenv("DB_PORT")
	dsn := databaseUsername + ":" + databasePassword + "@tcp(" + databaseAddress + ":" + databasePort + ")/" + databaseName + "?parseTime=true&charset=utf8mb4&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
}

func GetDatabase() *gorm.DB {
	return database
}
