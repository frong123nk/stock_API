package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	Db := connectDB()
	return Db
}

func connectDB() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Bangkok",
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connect", err)
	}
	return db
}
