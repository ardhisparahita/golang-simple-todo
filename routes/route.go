package routes

import (
	"golang-blog-api/handler"
	"golang-blog-api/middleware"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App, h *handler.UserHandler, t *handler.TodoHandler) {

	v1 := app.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)

	users := v1.Group("/users", middleware.JWTProtected())
	users.Get("/me", h.Me)
	users.Put("/:id", middleware.Ownership(), h.Update)
	users.Delete("/:id", middleware.Ownership(), h.Delete)
	users.Get("/:id", h.FindByID)
	users.Get("/", h.FindAll)

	todo := v1.Group("/todos", middleware.JWTProtected())
	todo.Post("/", t.Create)
	todo.Put("/:id", t.Update)
	todo.Delete("/:id", t.Delete)
	todo.Get("/:id", t.FindByID)
	todo.Get("/", t.FindByUserID)
}
