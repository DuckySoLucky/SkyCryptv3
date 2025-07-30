package routes

import (
	"fmt"
	"os"
	"path/filepath"
	"skycrypt/src/constants"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PotionHandlers(c *fiber.Ctx) error {
	timeNow := time.Now()
	potionType := c.Params("type")
	potionColor := c.Params("color")

	fmt.Printf("%s %s", potionType, potionColor)

	if potionType == "" || potionColor == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	imagePath := filepath.Join("cache", "potions", potionType, fmt.Sprintf("potion_%s_%s.png", potionType, potionColor))
	absPath, _ := filepath.Abs(imagePath)
	imageBytes, err := os.ReadFile(absPath)
	if err != nil {
		c.Status(404)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")
	fmt.Printf("Returning /api/potion/%s/%s in %s\n", potionType, potionColor, time.Since(timeNow))
	return c.Send(imageBytes)
}
