package src

import (
	"fmt"
	redis "skycrypt/src/db"
	"skycrypt/src/handlers"
	"skycrypt/src/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupApplication() error {
	err := redis.InitRedis("localhost:6379", "", 0)
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return nil
}

func SetupRoutes(app *fiber.App) {

	app.Get("/api/uuid/:username", routes.UUIDHandler)
	app.Get("/api/username/:uuid", routes.UsernameHandler)

	app.Get("/api/profiles/:uuid", routes.ProfilesHandler)
	app.Get("/api/player/:uuid", routes.PlayerHandler)
	app.Get("/api/museum/:profileId", routes.MuseumHandler)

	app.Get("/api/stats/:uuid/:profileId", routes.StatsHandler)
	app.Get("/api/stats/:uuid", routes.StatsHandler)

	// Root endpoint
	app.Get("/", handlers.HelloHandler)
}
