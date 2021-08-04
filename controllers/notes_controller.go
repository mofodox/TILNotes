package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/database"
	"github.com/mofodox/TILNotes/models"
)

// GET – retrieve all notes
func GetAllNotes(ctx *fiber.Ctx) error {
	var notes []models.Note

	database.DB.Model(&models.Note{}).Find(&notes)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": notes,
	})
}

// POST – add a note
func AddNote(ctx *fiber.Ctx) error {
	var note models.Note

	if err := ctx.BodyParser(&note); err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Create(&note)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note successfully created",
		"data": note,
	})
}
