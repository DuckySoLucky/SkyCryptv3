package routes

import (
	"fmt"
	"skycrypt/src/constants"
	"skycrypt/src/lib"

	"github.com/gofiber/fiber/v2"
)

func LeatherHandlers(c *fiber.Ctx) error {
	// timeNow := time.Now()
	armorType := c.Params("type")
	armorColor := c.Params("color")
	if armorType == "" || armorColor == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	imageBytes, err := lib.RenderArmor(armorType, armorColor)
	if err != nil {
		fmt.Printf("Error rendering armor: %v\n", err)
		c.Status(404)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")
	// fmt.Printf("Returning /api/leather/%s/%s in %s\n", armorType, armorColor, time.Since(timeNow))
	return c.Send(imageBytes)
}
