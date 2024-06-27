package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"testTask/cmd/internal/db/models"
)

type IHandlerService interface {
	GetAllUsersProfiles(ctx context.Context) ([]models.ProfileAPI, error)
	GetUsersProfilesByUsername(ctx context.Context, username string) ([]models.ProfileAPI, error)
}

type Handler struct {
	s IHandlerService
}

func NewHandler(s IHandlerService) *Handler {
	return &Handler{s: s}
}

func (h *Handler) UserProfiles(c *fiber.Ctx) error {
	var (
		res []models.ProfileAPI
		err error
	)

	param := c.Query("username")
	if param != "" {
		res, err = h.s.GetUsersProfilesByUsername(c.Context(), param)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	} else {
		res, err = h.s.GetAllUsersProfiles(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"status": "ok", "data": res})
}
