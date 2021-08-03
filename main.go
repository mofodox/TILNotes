package main

import (
	"fmt"
	"log"

	"github.com/mofodox/TILNotes/database"

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

	// Create the connection to our database
	if err := database.Connect(); err != nil {
		log.Printf("Unable to connect the database: %v\n", err.Error())
		panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	app.Get("/", hello)

	app.Listen(":1337")
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}