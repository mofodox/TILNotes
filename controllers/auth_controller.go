package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/database"
	"github.com/mofodox/TILNotes/models"
	"github.com/mofodox/TILNotes/util"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// Register POST - register user
func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to hash the password due to server error",
		})
	}

	user := models.User{
		Email: data["email"],
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data": user,
	})
}

// Login Post - login user
func Login(ctx *fiber.Ctx) error {
	var (
		data map[string]string
		user models.User
	)

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to process JSON request",
		})
	}

	database.DB.Model(&user).Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User ID %v not found", user.ID),
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to create token due to server error",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"jwt-token": token,
		"data": user,
	})
}

// CurrentUser GET - retrieve current logged in user
func CurrentUser(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	userId, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Preload("Notes.Category").Preload("Notes").Where("id = ?", userId).First(&user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"jwt-token": cookie,
		"data": user,
	})
}

// Logout POST - logout user
func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successfully",
	})
}