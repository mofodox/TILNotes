package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New()
	
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", hello)

	app.Listen(":1337")
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}