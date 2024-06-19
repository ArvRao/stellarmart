package database

import (
	"log"
	"os"

	I "github.com/ArvRao/ecommerce-app/interfaces"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabaseConnection() I.IDatabase {
	return Database{}
}

func (Database) CloseDb(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}

func (Database) OpenDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DNS")), &gorm.Config{})
	if err != nil {

		log.Fatal("error in connecting database : ", err)
	}
	return db
}
