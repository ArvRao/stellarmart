// cmd/yourapp/main.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArvRao/ecommerce-app/internal/database"
	"github.com/ArvRao/ecommerce-app/internal/routes"
	_ "github.com/lib/pq"
)

func main() {
	app := routes.SetupRoutes()

	database.DbConnection()

	// Start the Go Fiber app
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Sprintln("Starting your app on ", port)
	if err := app.Listen(host + ":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
