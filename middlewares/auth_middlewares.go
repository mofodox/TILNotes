package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/TILNotes/util"
)

func AuthRequired(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return ctx.Next()
}
