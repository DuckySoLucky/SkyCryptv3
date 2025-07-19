package main

import (
	"skycrypt/src"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:                   true,  // Fork processes for max CPU utilization
		ServerHeader:              "",    // Remove server header for slight perf gain
		DisableKeepalive:          false, // Keep connections alive
		DisableDefaultDate:        true,  // Disable date header
		DisableDefaultContentType: false,
		BodyLimit:                 10 << 20, // 10MB
		ReadBufferSize:            4096,
		WriteBufferSize:           4096,
		ReadTimeout:               0, // No timeout for max throughput
		WriteTimeout:              0,
		IdleTimeout:               0,
	})

	app.Use(recover.New())
	app.Use(cors.New())

	err := src.SetupApplication()
	if err != nil {
		panic(err)
	}

	src.SetupRoutes(app)

	app.Listen(":8080")
}
