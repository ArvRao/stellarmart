// cmd/yourapp/main.go

package main

import (
	"log"

	"github.com/ArvRao/ecommerce-app/internal/app/routes"
	"github.com/ArvRao/ecommerce-app/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Configure routes using Go Fiber
	app := routes.SetupRoutes()

	// database connection
	database.DbConnection()

	// Start the Go Fiber app
	log.Println("Starting your app on :8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
