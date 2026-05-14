package main

import (
	database "golang-blog-api/app"
	"golang-blog-api/config"
	"golang-blog-api/exception"
	"golang-blog-api/handler"
	"golang-blog-api/repository"
	"golang-blog-api/routes"
	"golang-blog-api/service"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadConfig()

	database.ConnectDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	})

	repo := repository.NewUserRepository(database.DB)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	routes.Setup(app, h)

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{}))
}
