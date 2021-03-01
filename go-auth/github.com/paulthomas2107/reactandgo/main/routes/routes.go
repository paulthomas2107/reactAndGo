package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulthomas2107/reactandgo/github.com/paulthomas2107/reactandgo/main/controllers"
)

// Setup ...
func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
