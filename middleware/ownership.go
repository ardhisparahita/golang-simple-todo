package middleware

import "github.com/gofiber/fiber/v3"

func Ownership() fiber.Handler {
	return func(c fiber.Ctx) error {
		userId, ok := c.Locals("userId").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		paramId := c.Params("id")

		if userId != paramId {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden",
			})
		}

		return c.Next()
	}
}
