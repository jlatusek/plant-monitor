package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"plant_monitor/configuration"
	"plant_monitor/controllers"
	"plant_monitor/database"
)

func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func main() {
	configuration.ServerConfiguration.Load(".env")
	configuration.ServerConfiguration.Print()
	database.MI.Connect()
	database.MI.ListDatabases()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	api := app.Group("/api", middleware)
	v1 := api.Group("/v1", middleware)

	plant := v1.Group("plant", middleware)
	plant.Get("/", controllers.GetPlantAll)
	plant.Get("/:id", controllers.GetPlant)

	err := app.Listen(":3000")
	if err != nil {
		log.Println(err)
	}
	database.MI.Disconnect()
}

