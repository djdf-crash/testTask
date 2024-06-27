package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type IAuthMiddlewareService interface {
	IsExistAuthKey(ctx context.Context, key string) bool
}

const headerKey = "Api-Key"

func NewAuthMiddleware(s IAuthMiddlewareService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get(headerKey)
		if key == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}
		exist := s.IsExistAuthKey(c.Context(), key)
		if !exist {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
		}

		return c.Next()
	}
}
