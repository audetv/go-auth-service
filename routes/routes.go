package routes

import (
	"github.com/audetv/go-auth-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
}
