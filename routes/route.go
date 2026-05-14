package routes

import (
	"golang-blog-api/handler"
	"golang-blog-api/middleware"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App, h *handler.UserHandler) {

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
}
