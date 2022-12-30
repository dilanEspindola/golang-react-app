package config

import (
	"log"
	"os"

	"anime-hentai-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInstace struct {
	Db *gorm.DB
}

var Database DbInstace

func DbConnection() {
	dsn := "root:dilan@tcp(localhost:3307)/fullstackapp_go_react?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connecting db", err)
		os.Exit(2)
	}

	log.Println("db is connected")

	log.Println("Running migrations")
	db.AutoMigrate(&models.User{})

	Database = DbInstace{Db: db}
}
