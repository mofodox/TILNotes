package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/controllers"
)

func SetupRoutes(app *fiber.App) {
	noteApiV1 := app.Group("/api/v1/notes")

	noteApiV1.Get("/", controllers.GetAllNotes)
	noteApiV1.Post("/", controllers.AddNote)
	noteApiV1.Get("/:id", controllers.GetNote)
	noteApiV1.Put("/edit/:id", controllers.EditNote)
	noteApiV1.Delete("/delete/:id", controllers.DeleteNote)
}
