package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"plant_monitor/configuration"
	"plant_monitor/controllers"
	"plant_monitor/database"
)

func middleware (c *fiber.Ctx) error {
	fmt.Println("Middleware called")
	return c.Next()
}

func main() {
	configuration.ServerConfiguration.Load()
	configuration.ServerConfiguration.Print()
	database.MI.Connect()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	api := app.Group("/api", middleware)
	v1 := api.Group("/v1", middleware)

	plant := v1.Group("plant", middleware)
	plant.Get("/:id", controllers.GetPlant)

	app.Listen(":3000")
	database.MI.Disconnect()
}

