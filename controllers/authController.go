package controllers

import (
	"salesdashboard/database"
	"salesdashboard/utils"
	"salesdashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string] string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		RoleId: 1 ,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)
		return c.JSON(user)
	}


	func Login(c *fiber.Ctx) error {

		var data map[string]string
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		var user models.User
		database.DB.Where("email = ?", data["email"]).First(&user)

		if user.Id == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "not found",
			})
		}

		if err := user.ComparePassword(data["password"]); err != nil {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{
				"message": "Incorrect password",
			})
		}

		token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		cookie := fiber.Cookie{
			Name: "jwt",
			Value: token,
			Expires: time.Now().Add(time.Hour*24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwtCookie(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}