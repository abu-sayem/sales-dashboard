package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}
	fileName := ""
	files := form.File["image"]

	for _, file := range files {
		fileName = file.Filename

		if err := c.SaveFile(file, "./uploads/" + fileName); err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + fileName,
		"fileName": fileName,
	})
}