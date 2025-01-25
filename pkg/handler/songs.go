package handler

import (
	"net/http"
	"strconv"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddSongResponse struct {
	SongId uint `json:"song_id"`
}

// @Summary		Add a Song
// @Security		ApiKeyAuth
// @Tags			Songs
// @Description	Add a song to user's collection
// @Accept			json
// @Produce		json
// @Param			input		body		musiclib.Song	true	"Song details"
// @Success		200			{object}	AddSongResponse	"Successfully added"
// @Failure		400,401,404	{object}	MyError			"4** error"
// @Failure		500			{object}	MyError			"5** error"
// @Router			/api/songs [post]
func (h *Handler) addUserSong(ctx *fiber.Ctx) error {
	var songInput musiclib.Song

	if err := ctx.BodyParser(&songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	if err := validate.Struct(songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("type assertation failed")
		return ctx.Status(http.StatusUnauthorized).JSON(MyError{
			Err: "invalid user id",
		})
	}
	songId, err := h.services.Song.AddUserSong(userId, songInput)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	logrus.Printf("song added successfully, song id =%v", songId)
	return ctx.Status(http.StatusOK).JSON(AddSongResponse{
		SongId: songId,
	})
}

// @Summary		Get user songs
// @Security		ApiKeyAuth
// @Tags			Songs
// @Description	Retrieve user songs with pagination
// @Accept			json
// @Produce		json
// @Param			page		query		int				false	"Page number (default: 1)"
// @Param			pageSize	query		int				false	"Number of items per page (default: 10)"
// @Success		200			{array}		musiclib.Song	"List of songs"
// @Failure		400			{object}	MyError			"Invalid query parameters"
// @Failure		401			{object}	MyError			"Unauthorized"
// @Failure		500			{object}	MyError			"Internal server error"
// @Router			/api/songs [get]
func (h *Handler) getUserSongs(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)

	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(MyError{
			Err: "invalid user id",
		})
	}

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		logrus.Error("invalid page parameter")
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: "invalid page parametr",
		})
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		logrus.Error("invalid pageSize parameter")
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: "invalig pageSize parametr",
		})
	}

	songs, err := h.services.Song.GetUserSongs(userId, page, pageSize)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}
	if songs == nil {
		songs = []musiclib.Song{}
	}

	return ctx.Status(http.StatusOK).JSON(songs)
}

// @Summary		Get user songs by ID
// @Security		ApiKeyAuth
// @Tags			Songs
// @Description	Retrieve user songs by ID
// @Accept			json
// @Produce		json
// @Param			id	path		int				true	"Song ID"
// @Success		200	{object}	musiclib.Song	"Song"
// @Failure		400	{object}	MyError			"Invalid query parameters"
// @Failure		401	{object}	MyError			"Unauthorized"
// @Failure		500	{object}	MyError			"Internal server error"
// @Router			/api/songs/{id} [get]
func (h *Handler) getUserSongById(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(MyError{
			Err: "invalid user id",
		})
	}

	songId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	song, err := h.services.Song.GetUserSongById(userId, songId)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(song)
}

// @Summary		Update user song by ID
// @Security		ApiKeyAuth
// @Tags			Songs
// @Description	Updating user song info by ID
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Song ID"
// @Param			input	body		musiclib.Song	true	"Song data to update"
// @Success		200		{object}	musiclib.Song	"Updated Song"
// @Failure		400		{object}	MyError			"Invalid query parameters"
// @Failure		401		{object}	MyError			"Unauthorized"
// @Failure		500		{object}	MyError			"Internal server error"
// @Router			/api/songs/{id} [put]
func (h *Handler) updateUserSongInfo(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("invalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(MyError{
			Err: "invalid user id",
		})
	}

	songId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	var songInput musiclib.Song
	if err := ctx.BodyParser(&songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	if err := validate.Struct(songInput); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	updatedSong, err := h.services.Song.UpdateUserSongInfo(userId, songId, songInput)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	logrus.Printf("song %s is updated successfully", updatedSong.Title)
	return ctx.Status(http.StatusOK).JSON(updatedSong)
}

type DeleteSongResponce struct {
	Status string `json:"status"`
}

// @Summary		Delete user song by ID
// @Security		ApiKeyAuth
// @Tags			Songs
// @Description	Deleting user song by ID
// @Accept			json
// @Produce		json
// @Param			id	path		int					true	"Song ID"
// @Success		200	{object}	DeleteSongResponce	"Deletion status"
// @Example		200 {object} DeleteSongResponce {"status": "ok"}
// @Failure		400	{object}	MyError	"Invalid query parameters"
// @Failure		401	{object}	MyError	"Unauthorized"
// @Failure		500	{object}	MyError	"Internal server error"
// @Router			/api/songs/{id} [delete]
func (h *Handler) deleteUserSongById(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(uint)
	if !ok {
		logrus.Error("ivalid user id")
		return ctx.Status(http.StatusUnauthorized).JSON(MyError{
			Err: "invalid user id",
		})
	}

	songId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	if err := h.services.Song.DeleteUserSongByID(userId, songId); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	logrus.Printf("song succesfully deleted songId = %v userId = %v\n", songId, userId)
	return ctx.Status(http.StatusOK).JSON(DeleteSongResponce{
		Status: "ok",
	})
}
