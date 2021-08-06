package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/database"
	"github.com/mofodox/TILNotes/models"
)

// GetAllNotes GET – retrieve all notes
func GetAllNotes(ctx *fiber.Ctx) error {
	var notes []models.Note

	database.DB.Model(&models.Note{}).Preload("Category").Find(&notes)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": notes,
	})
}

// AddNote POST – add a note
func AddNote(ctx *fiber.Ctx) error {
	var note models.Note

	if err := ctx.BodyParser(&note); err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Create(&note)

	if note.ID != 0 {
		err := database.DB.Model(&models.Category{}).Where("id = ?", note.CategoryId).Take(&note.Category).Error
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Create note: server error",
			})
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note successfully created",
		"data": note,
	})
}

// GetNote GET - retrieve a note
func GetNote(ctx *fiber.Ctx) error {
	var note models.Note

	noteId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Error converting the noteId to int64",
		})
	}

	// If the result (noteId != note.ID) does not exist, return server error code and message
	if result := database.DB.Model(&models.Note{}).Preload("Category").Where("id = ?", noteId).First(&note); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to retrieve the specified noteId due to server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": note,
	})
}

// EditNote PUT - update a note
func EditNote(ctx *fiber.Ctx) error {
	noteId, _ := strconv.ParseUint(ctx.Params("id"), 10, 64)

	note := models.Note{
		ID: noteId,
	}

	if err := ctx.BodyParser(&note); err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Model(&note).Updates(note)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Note ID %v has been updated", noteId),
		"data": note,
	})
}

// DeleteNote DELETE - remove a single note based on note id
func DeleteNote(ctx *fiber.Ctx) error {
	var note models.Note

	noteId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	// Get the note to delete from our database
	db := database.DB.Model(&models.Note{}).Where("id = ?", noteId).Find(&note)
	// Delete the note from our database
	db.Delete(&note)

	if db.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to perform DELETE request due to server error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": fmt.Sprintf("Note ID %v has been deleted", noteId),
	})
}
