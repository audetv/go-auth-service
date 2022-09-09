package main

import (
	"github.com/audetv/go-auth-service/database"
	"github.com/audetv/go-auth-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
