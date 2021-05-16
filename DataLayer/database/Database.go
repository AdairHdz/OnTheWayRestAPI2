package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func init() {
	dsn := "adair:adahplf0015@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("Could not connect to database")
	}
}

func GetDatabase() *gorm.DB{
	return database
}