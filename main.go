// cmd/yourapp/main.go

package main

import (
	"log"

	"github.com/ArvRao/ecommerce-app/internal/app/routes"
)

func main() {
	// Configure routes using Go Fiber
	app := routes.SetupRoutes()

	// authentication

	// Start the Go Fiber app
	log.Println("Starting your app on :8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
