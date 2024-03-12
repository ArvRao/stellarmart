// database/database.go

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "StellarCart"
)

// NewDB creates a new PostgreSQL database connection.
func DbConnection() {
	// database connection
	psqlConn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	// authentication

	// Connect to database
	db, err := sql.Open("postgres", psqlConn)
	checkError(err)
	defer db.Close()

	insertStmt := InsertCommand()
	_, e := db.Exec(insertStmt)

	checkError(e)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
