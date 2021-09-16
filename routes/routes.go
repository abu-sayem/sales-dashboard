package routes

import (
	"salesdashboard/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("api/logout", controllers.Logout)
	app.Get("api/other", controllers.Other)
	app.Get("api/user", controllers.User)
}
