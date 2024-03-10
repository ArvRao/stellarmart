// cmd/yourapp/main.go

package main

import (
	"fmt"
	"log"

	"github.com/ArvRao/ecommerce-app/internal/app/routes"
	"github.com/supabase-community/supabase-go"
)

func main() {
	// Configure routes using Go Fiber
	app := routes.SetupRoutes()

	// authentication

	client, err := supabase.NewClient("https://jexftzhwzotwcyceslfz.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpleGZ0emh3em90d2N5Y2VzbGZ6Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTAwOTMyMjUsImV4cCI6MjAyNTY2OTIyNX0.6z1yxnUkWBB5SaJRZJuuCZM-rns1cSc6jOlkCL_kTT8", nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, _ := client.From("countries").Select("*", "exact", false).Execute()
	// Start the Go Fiber app
	log.Println("Starting your app on :8080...")
	log.Println(data, count)
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
