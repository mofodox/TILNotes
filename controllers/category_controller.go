package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/database"
	"github.com/mofodox/TILNotes/models"
)

// GET – retrieve all categories
func GetAllCategories(ctx *fiber.Ctx) error {
	var categories []models.Category

	database.DB.Model(&models.Category{}).Find(&categories)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": categories,
	})
}

// POST – add a note
func AddCategory(ctx *fiber.Ctx) error {
	var category models.Category

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Create(&category)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note successfully created",
		"data": category,
	})
}

// GET - retrieve a note
func GetCategory(ctx *fiber.Ctx) error {
	var category models.Category

	categoryId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Error converting the noteId to int64",
		})
	}

	// If the result (noteId != note.ID) does not exist, return server error code and message
	if result := database.DB.Model(&models.Category{}).Where("id = ?", categoryId).First(&category); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to retrieve the specified noteId due to server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": category,
	})
}

// PUT - update a note
func EditCategory(ctx *fiber.Ctx) error {
	categoryId, _ := strconv.ParseUint(ctx.Params("id"), 10, 64)

	category := models.Note{
		ID: categoryId,
	}

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Model(&category).Updates(category)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Note ID %v has been updated", categoryId),
		"data": category,
	})
}

// DELETE - remove a single note based on note id
func DeleteCategory(ctx *fiber.Ctx) error {
	var category models.Category

	categoryId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	// Get the note to delete from our database
	db := database.DB.Model(&models.Category{}).Where("id = ?", categoryId).Find(&category)
	// Delete the note from our database
	db.Delete(&category)

	if db.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to perform DELETE request due to server error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": fmt.Sprintf("Note ID %v has been deleted", categoryId),
	})
}
