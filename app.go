package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ArvRao/ecommerce-app/internal/users/controllers"
	repository "github.com/ArvRao/ecommerce-app/internal/users/repositories"
	service "github.com/ArvRao/ecommerce-app/internal/users/services"
	"github.com/ArvRao/ecommerce-app/pkg/database"
	"github.com/ArvRao/ecommerce-app/pkg/routes"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitApp() {
	// Get the absolute path to the root directory
	root, err := filepath.Abs("../")
	if err != nil {
		log.Fatalf("Error getting root directory: %v", err)
	}

	// Construct the path to the .env file
	envPath := filepath.Join(root, ".env")

	// Load the .env file from the specified path
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database.DbConnection()
	// app := routes.SetupRoutes()
	validate := validator.New()

	// init repo, service, controller
	userRepository := repository.NewUserRepositoryImpl(database.DB.Db)

	userService := service.NewUserServiceImpl(userRepository, validate)

	userController := controllers.NewUserController(userService)
	// update here to point to router -> router.go file
	app := routes.NewRouter(userController)
	// routes := routes.NewRouter(userController)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Status":  "success",
			"message": "welcome to stellarmart",
		})
	})
	// Start the Go Fiber app
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Sprintln("Starting your app on ", port)
	if err := app.Listen(host + ":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
