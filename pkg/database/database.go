// database/database.go

package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/models"
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	// Construct the path to the .env file
	envPath := filepath.Join(root, ".env")

	// Load the .env file from the specified path
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func DbConnection() {
	// loadEnv()
	psqlConn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable TimeZone=Asia/Kolkata", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Connect to database
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.ErrorPanic(err)
	log.Println("Connected successfully to the database!")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&Users.User{}, &Orders.Order{})
	DB = DBInstance{
		Db: db,
	}
}
