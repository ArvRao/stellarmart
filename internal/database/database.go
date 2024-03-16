// database/database.go

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func DbConnection() {
	loadEnv()
	psqlConn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable TimeZone=Asia/Kolkata", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Connect to database
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.ErrorPanic(err)
	log.Println("Connected successfully to the database!")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.User{})
	DB = DBInstance{
		Db: db,
	}
}
