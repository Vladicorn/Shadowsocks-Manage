package rout

import (
	"sshodows/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")
	api.Get("/create/:port", controllers.CreateSS)

	//	api.Get("/api/test/user", controllers.Test)
	//adminAuth := api.Use(middlewares.Auth)
}
