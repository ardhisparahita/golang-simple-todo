package middleware

import "github.com/gofiber/fiber/v3"

func Ownership() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := c.Locals("userId").(string)

		paramId := c.Params("id")

		if userId != paramId {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden",
			})
		}

		return c.Next()
	}
}
