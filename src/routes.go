package src

import (
	"fmt"
	"log"
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
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
		log.Fatal("error loading .env file")
	}

	if err := notenoughupdates.InitializeNEURepository(); err != nil {
		return fmt.Errorf("error initializing repository: %v", err)
	}

	if err := notenoughupdates.UpdateNEURepository(); err != nil {
		return fmt.Errorf("error updating repository: %v", err)
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

	app.Get("/", handlers.HelloHandler)
}
