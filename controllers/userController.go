package controllers

import (
	"salesdashboard/database"
	"salesdashboard/models"

	"github.com/gofiber/fiber/v2"
)

func AllUser(c *fiber.Ctx) error {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(users)
	}


func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}