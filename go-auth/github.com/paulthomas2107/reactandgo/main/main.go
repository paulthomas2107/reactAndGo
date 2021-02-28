package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulthomas2107/reactandgo/github.com/paulthomas2107/reactandgo/main/database"
	"github.com/paulthomas2107/reactandgo/github.com/paulthomas2107/reactandgo/main/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
