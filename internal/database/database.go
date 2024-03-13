// database/database.go

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func DbConnection() {
	loadEnv()

	// database connection
	psqlConn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Connect to database
	db, err := sql.Open("postgres", psqlConn)
	checkError(err)
	defer db.Close()

	insertStmt := InsertUser()
	_, e := db.Exec(insertStmt)

	checkError(e)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
