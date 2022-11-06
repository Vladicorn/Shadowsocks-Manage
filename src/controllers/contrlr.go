package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sshodows/src/shadows"
)

func CreateSS(c *fiber.Ctx) error {
	port := c.Params("port")
	log.Println("Создание сервера SS,  ss://AEAD_CHACHA20_POLY1305:password@:", port)

	go shadows.Connect(port)
	return c.JSON(fiber.Map{"status": "success"})
}
