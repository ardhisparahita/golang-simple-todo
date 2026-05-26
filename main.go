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

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	todoRepo := repository.NewTodoRepository(database.DB)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	routes.Setup(app, userHandler, todoHandler)

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{}))
}
