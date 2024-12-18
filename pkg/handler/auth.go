package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) signIn(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "sign-in",
	})
	return nil
}

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "sign-up",
	})
	return nil
}
