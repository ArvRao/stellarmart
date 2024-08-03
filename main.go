package main

import (
	"time"

	"github.com/ArvRao/ecommerce-app/database"
	"github.com/ArvRao/ecommerce-app/routes"
	utils "github.com/ArvRao/ecommerce-app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	time.Sleep(time.Second * 7)
	database.SyncDatabase()

}

func main() {

	app := fiber.New()
	app.Use(logger.New())

	routes.AdminRoute(app)

	routes.UserRoute(app)

	utils.ListenAndShutdown(app)
}
