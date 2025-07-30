package routes

import (
	"bytes"
	"fmt"
	"image/png"
	"skycrypt/src/constants"
	"skycrypt/src/lib"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ItemHandlers(c *fiber.Ctx) error {
	timeNow := time.Now()
	textureId := c.Params("itemId")
	if textureId == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	texture, err := lib.RenderItem(textureId)
	if err != nil {
		c.Status(500)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")

	var buf bytes.Buffer
	if err := png.Encode(&buf, texture); err != nil {
		c.Status(500)
		return c.SendString("Failed to encode image")
	}

	fmt.Printf("Returning /api/head/%s in %s\n", textureId, time.Since(timeNow))

	return c.Send(buf.Bytes())
}
