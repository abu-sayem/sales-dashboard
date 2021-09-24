package controllers

import (
	"salesdashboard/database"
	"salesdashboard/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePermissions(c *fiber.Ctx) error {
	var data map[string] string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))
	permission := models.Permission{
		Id: uint(id),
		Name: data["name"],
	}

		database.DB.Create(&permission)

		return c.JSON(permission)
	}

func AllPermissions(c *fiber.Ctx) error {
		var permissions []models.Permission

		database.DB.Find(&permissions)

		return c.JSON(permissions)
	}