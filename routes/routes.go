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
	app.Get("api/users/:id", controllers.GeteUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CreatRole)
	app.Get("api/roles/:id", controllers.GeteRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)
	app.Post("api/permissions", controllers.CreatePermissions)
}
