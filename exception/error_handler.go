package exception

import (
	"errors"
	"golang-blog-api/helper"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	response := helper.WebResponse{
		StatusCode: code,
		Status:     "Error",
		Message:    err.Error(),
	}

	var notFoundError *NotFoundError
	if errors.As(err, &notFoundError) {
		code = fiber.StatusNotFound
		response.StatusCode = code
		response.Message = notFoundError.Message
	}

	return c.Status(code).JSON(response)
}
