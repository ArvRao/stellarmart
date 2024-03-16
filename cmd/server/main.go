package server

// cmd/yourapp/main.go

import (
	"fmt"
	"log"
	"os"

	"github.com/ArvRao/ecommerce-app/internal/controllers"
	"github.com/ArvRao/ecommerce-app/internal/database"
	"github.com/ArvRao/ecommerce-app/internal/repository"
	"github.com/ArvRao/ecommerce-app/internal/routes"
	"github.com/ArvRao/ecommerce-app/internal/service"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	database.DbConnection()
	// app := routes.SetupRoutes()
	validate := validator.New()

	// init repo, service, controller
	userRepository := repository.NewUserRepositoryImpl(database.DB.Db)

	userService := service.NewUserServiceImpl(userRepository, validate)

	userController := controllers.NewUserController(userService)

	routes := routes.NewRouter(userController)
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Status":  "success",
			"message": "welcome to stellarmart",
		})
	})
	app.Mount("/api", routes)
	// Start the Go Fiber app
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Sprintln("Starting your app on ", port)
	if err := app.Listen(host + ":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
