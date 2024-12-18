package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) userAuthMiddleware(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "userAuthMiddleware",
	})
	return nil
}
