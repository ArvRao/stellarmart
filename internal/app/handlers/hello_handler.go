// internal/app/handlers/hello_handler.go

package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HelloHandler handles requests for the "/hello" endpoint using Go Fiber.
func HelloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
