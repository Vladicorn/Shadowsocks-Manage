package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"sshodows/src/route"
)

var IP string = "http://127.0.0.1,https://127.0.0.1"

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     IP,
		AllowHeaders:     "Origin, Content-Type, Accept,Access-Control-Allow-Origin",
	}))

	rout.Setup(app)
	log.Println("перейдите на http://127.0.0.1:8086/api/create/:port ,где port - порт подключения")
	app.Listen(":8086")

}
