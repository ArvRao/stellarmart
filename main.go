package main

import (
	"time"

	"github.com/ArvRao/ecommerce-app/database"
	"github.com/ArvRao/ecommerce-app/routes"
	utils "github.com/ArvRao/ecommerce-app/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	time.Sleep(time.Second * 7)
	database.SyncDatabase()

}

func main() {

	app := fiber.New()

	routes.AdminRoute(app)

	routes.UserRoute(app)

	utils.ListenAndShutdown(app)
}
