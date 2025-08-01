package routes

import (
	"skycrypt/src/constants"
	"skycrypt/src/lib"

	"github.com/gofiber/fiber/v2"
)

func HeadHandlers(c *fiber.Ctx) error {
	// timeNow := time.Now()
	textureId := c.Params("textureId")
	if textureId == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	textureBytes := lib.RenderHead(textureId)
	if textureBytes == nil {
		c.Status(500)
		return c.SendString("Failed to render head")
	}

	c.Type("png")
	// fmt.Printf("Returning /api/head/%s in %s\n", textureId, time.Since(timeNow))
	return c.Send(textureBytes)
}
