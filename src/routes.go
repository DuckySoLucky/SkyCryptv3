package src

import (
	"fmt"
	"log"
	"os"
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/handlers"
	"skycrypt/src/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
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

	if err := api.LoadSkyBlockItems(); err != nil {
		return fmt.Errorf("error loading SkyBlock items: %v", err)
	}

	if err := notenoughupdates.InitializeNEURepository(); err != nil {
		return fmt.Errorf("error initializing repository: %v", err)
	}

	if err := notenoughupdates.UpdateNEURepository(); err != nil {
		return fmt.Errorf("error updating repository: %v", err)
	}

	err = notenoughupdates.ParseNEURepository()
	if err != nil {
		return fmt.Errorf("error parsing NEU repository: %v", err)
	}

	return nil
}

func SetupRoutes(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	if os.Getenv("DEV") == "false" {
		fmt.Println("[ENVIROMENT] Running in production mode")
		app.Use(etag.New())
		app.Use("/api", cache.New(cache.Config{
			Expiration:   5 * time.Minute,
			CacheControl: true,
		}))
	}

	api := app.Group("/api")

	api.Get("/uuid/:username", routes.UUIDHandler)
	api.Get("/username/:uuid", routes.UsernameHandler)

	api.Get("/profiles/:uuid", routes.ProfilesHandler)
	api.Get("/player/:uuid", routes.PlayerHandler)
	api.Get("/museum/:profileId", routes.MuseumHandler)

	api.Get("/stats/:uuid/:profileId", routes.StatsHandler)
	api.Get("/stats/:uuid", routes.StatsHandler)

	api.Get("/gear/:uuid/:profileId", routes.GearHandler)
	api.Get("/accessories/:uuid/:profileId", routes.AccessoriesHandler)
	api.Get("/pets/:uuid/:profileId", routes.PetsHandler)

	api.Get("/inventory/:uuid/:profileId/:inventoryId", routes.InventoryHandler)
	api.Get("/inventory/:uuid/:profileId/:inventoryId/:search", routes.InventoryHandler)

	api.Get("/skills/:uuid/:profileId", routes.SkillsHandler)

	// Root route
	app.Get("/", handlers.HelloHandler)
}
