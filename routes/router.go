package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/controllers"
)

func SetupRoutes(app *fiber.App) {
	noteApiV1 := app.Group("/api/v1/notes")
	categoryApiV1 := app.Group("/api/v1/categories")
	userAuthApiV1 := app.Group("/api/v1/users/auth")

	userAuthApiV1.Post("/register", controllers.Register)
	userAuthApiV1.Post("/login", controllers.Login)
	userAuthApiV1.Get("/current_user", controllers.CurrentUser)
	userAuthApiV1.Post("/logout", controllers.Logout)

	noteApiV1.Get("/", controllers.GetAllNotes)
	noteApiV1.Post("/", controllers.AddNote)
	noteApiV1.Get("/:id", controllers.GetNote)
	noteApiV1.Put("/edit/:id", controllers.EditNote)
	noteApiV1.Delete("/delete/:id", controllers.DeleteNote)

	categoryApiV1.Get("/", controllers.GetAllCategories)
	categoryApiV1.Post("/", controllers.AddCategory)
	categoryApiV1.Get("/:id", controllers.GetCategory)
	categoryApiV1.Put("/edit/:id", controllers.EditCategory)
	categoryApiV1.Delete("/delete/:id", controllers.DeleteCategory)
}
