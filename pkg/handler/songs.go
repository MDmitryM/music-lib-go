package handler

import (
	"net/http"
	"strconv"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addUserSong(ctx *fiber.Ctx) error {
	var songInput musiclib.Song

	if err := ctx.BodyParser(&songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("type assertation failed")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid uesr id",
		})
	}
	songId, err := h.services.Song.AddUserSong(userId, songInput)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.Printf("song added successfully, song id =%v", songId)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"song_id": songId,
	})
}

func (h *Handler) getUserSongs(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)

	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		logrus.Error("invalid page parameter")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid page parameter",
		})
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		logrus.Error("invalid pageSize parameter")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid pageSize parameter",
		})
	}

	songs, err := h.services.Song.GetUserSongs(userId, page, pageSize)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(songs)
}

func (h *Handler) getUserSongById(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	songId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	song, err := h.services.Song.GetUserSongById(userId, songId)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(song)
}

func (h *Handler) updateUserSongInfo(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	songId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var songInput musiclib.Song
	if err := ctx.BodyParser(&songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedSong, err := h.services.Song.UpdateUserSongInfo(userId, songId, songInput)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.Printf("song %s is updated successfully", updatedSong.Title)
	return ctx.Status(http.StatusOK).JSON(updatedSong)
}

func (h *Handler) deleteUserSongById(ctx *fiber.Ctx) error {
	logrus.Println(ctx.Locals("user_id"))
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "deleteUserSongById",
	})
	return nil
}
