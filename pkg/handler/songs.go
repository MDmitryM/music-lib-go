package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addUserSong(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "addUserSong",
	})
	return nil
}

func (h *Handler) getUserSongs(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getUserSongs",
	})
	return nil
}

func (h *Handler) getUserSongById(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getUserSongById",
	})
	return nil
}

func (h *Handler) updateUserSongInfo(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "updateUserSongInfo",
	})
	return nil
}

func (h *Handler) deleteUserSongById(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "deleteUserSongById",
	})
	return nil
}
