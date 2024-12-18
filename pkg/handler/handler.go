package handler

import (
	"github.com/gofiber/fiber"
)

type Handler struct {
}

func (h *Handler) SetupRouts() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth") // /auth
	auth.Post("/sign-up")      // /auth/sign-in
	auth.Post("/sign-in")      // /auth/sign-up

	api := app.Group("/api") // /api

	songs := api.Group("/songs") // /api/songs
	songs.Post("/")              // /api/songs/ add new song
	songs.Get("/")               // /api/songs/id get all songs
	songs.Get("/:id")            // /api/songs/id get song by id
	songs.Put("/:id")            // /api/songs/id update songs info by id
	songs.Delete("/:id")         // /api/songs/id delete song by id

	return app
}
