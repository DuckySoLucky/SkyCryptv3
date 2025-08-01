package routes

import (
	"skycrypt/src/constants"
	"skycrypt/src/lib"

	"github.com/gofiber/fiber/v2"
)

func ItemHandlers(c *fiber.Ctx) error {
	// timeNow := time.Now()
	textureId := c.Params("itemId")
	if textureId == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	textureBytes, err := lib.RenderItem(textureId)
	if err != nil {
		c.Status(500)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")
	// fmt.Printf("Returning /api/item/%s in %s\n", textureId, time.Since(timeNow))
	return c.Send(textureBytes)
}
