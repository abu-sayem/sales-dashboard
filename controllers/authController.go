package controllers

import (
	"salesdashboard/database"
	"salesdashboard/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		Password: password,
 
	}
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

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{
				"message": "Incorrect password",
			})
		}
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer: strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
		})
		token, err := claims.SignedString(([]byte("secret")))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}


		return c.JSON(token)
	}