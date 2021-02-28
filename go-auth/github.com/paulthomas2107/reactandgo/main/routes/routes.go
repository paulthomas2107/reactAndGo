package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulthomas2107/reactandgo/github.com/paulthomas2107/reactandgo/main/controllers"
)

// Setup ...
func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
}