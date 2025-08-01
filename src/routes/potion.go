package routes

import (
	"fmt"
	"skycrypt/src/constants"
	"skycrypt/src/lib"

	"github.com/gofiber/fiber/v2"
)

func PotionHandlers(c *fiber.Ctx) error {
	// timeNow := time.Now()
	potionType := c.Params("type")
	potionColor := c.Params("color")
	if potionType == "" || potionColor == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	imageBytes, err := lib.RenderPotion(potionType, potionColor)
	if err != nil {
		fmt.Printf("Error rendering armor: %v\n", err)

		c.Status(404)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")
	// fmt.Printf("Returning /api/potion/%s/%s in %s\n", potionType, potionColor, time.Since(timeNow))
	return c.Send(imageBytes)
}
