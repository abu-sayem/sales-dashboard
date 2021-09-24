package controllers

import (
	"math"
	"salesdashboard/database"
	"salesdashboard/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProduct(c *fiber.Ctx) error {
		var product []models.Product

		limit := 1
		page, _ := strconv.Atoi(c.Query("page", "1"))
		offset := (page - 1) * limit
		var total int64


		database.DB.Offset(offset).Limit(limit).Find(&product)
		database.DB.Model(&models.Product{}).Count(&total)

		return c.JSON(fiber.Map{
			"data": product,
			"meta": fiber.Map{
				"total": total,
				"page": page,
				"last_page": math.Ceil(float64(total) / float64(limit)),
			},
		})
	}


func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GeteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product {
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil{
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product {
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}