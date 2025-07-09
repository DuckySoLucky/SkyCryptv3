package src

import (
	"fmt"
	"log"
	redis "skycrypt/src/db"
	"skycrypt/src/handlers"
	"skycrypt/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupApplication() error {
	err := redis.InitRedis("localhost:6379", "", 0)
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	app.Get("/api/inventory/:uuid", routes.InventoryHandler)
	app.Get("/api/inventory/:uuid/:profileId", routes.InventoryHandler)

	app.Get("/api/gear/:uuid/:profileId", routes.GearHandler)

	// Root endpoint
	app.Get("/", handlers.HelloHandler)
}
