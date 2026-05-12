package main

import (
	"golang-blog-api/config"
	database "golang-blog-api/db"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadConfig()

	database.ConnectDB()

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("server running")
	})

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{}))
}
