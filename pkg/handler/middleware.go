package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authHeader = "Authorization"
	userCtx    = "user_id"
)

func (h *Handler) userAuthMiddleware(ctx *fiber.Ctx) error {
	header := ctx.Get(authHeader)
	if header == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "authorization header is required",
		})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid auth header format",
		})
	}

	token := headerParts[1]
	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Locals(userCtx, userId)
	return ctx.Next()
}
