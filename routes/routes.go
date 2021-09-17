package routes

import (
	"salesdashboard/controllers"
	"salesdashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthnicated)

	app.Post("api/logout", controllers.Logout)
	app.Get("api/user", controllers.User)

	app.Get("api/users", controllers.AllUser)
	app.Post("api/users", controllers.CreateUser)
}
