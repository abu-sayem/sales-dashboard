package middleware

import (
	"salesdashboard/utils"

	"github.com/gofiber/fiber/v2"
)

func IsAuthnicated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, error := utils.ParseJwtCookie(cookie); error !=nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthnicated user",
		})
	}

	return c.Next()
}