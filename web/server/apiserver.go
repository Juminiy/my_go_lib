package server

import (
	"github.com/gofiber/fiber/v2"
)

func InitServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "alan-algo-api-server",
		AppName:       "Algo API Server",
	})
	// middleware.UseCsrf(app)
	return app
}
