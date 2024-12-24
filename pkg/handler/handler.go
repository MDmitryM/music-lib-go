package handler

import (
	"github.com/MDmitryM/music-lib-go/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) SetupRouts(app *fiber.App) {
	//Authorization group
	auth := app.Group("/auth")      // /auth
	auth.Post("/sign-up", h.signUp) // /auth/sign-in (registration)
	auth.Post("/sign-in", h.signIn) // /auth/sign-up (аутентификация)

	//API group (auth required)
	api := app.Group("/api", h.userAuthMiddleware) // /api

	//CRUD songs operations
	songs := api.Group("/songs")               // /api/songs
	songs.Post("", h.addUserSong)              // /api/songs add new song
	songs.Get("", h.getUserSongs)              // /api/songs get songs by pages
	songs.Get("/:id", h.getUserSongById)       // /api/songs/id get song by id
	songs.Put("/:id", h.updateUserSongInfo)    // /api/songs/id update songs info by id
	songs.Delete("/:id", h.deleteUserSongById) // /api/songs/id delete song by id
}
