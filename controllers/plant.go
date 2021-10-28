package controllers

import (
	"github.com/gofiber/fiber/v2"
	"plant_monitor/database"
	"plant_monitor/models"
)

func GetPlant(c *fiber.Ctx) error {
	plant, err := models.GetPlantById(database.MI, c.Params("id"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(plant)
}
func GetPlantAll(c *fiber.Ctx) error {
	plants, err := models.GetPlantAll(database.MI)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(plants)
}
