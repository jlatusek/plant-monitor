package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetPlant(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Plant id: %s", c.Params("id"))
	return c.SendString(msg)
}
