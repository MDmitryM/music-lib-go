package handler

import (
	"net/http"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponce struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"access_token"`
}

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var input musiclib.User

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.Printf("successfully signed up, userId=%v", id)
	return ctx.Status(http.StatusOK).JSON(RegisterResponce{
		ID:          id,
		AccessToken: token,
	})
}

func (h *Handler) signIn(ctx *fiber.Ctx) error {
	var input SignInInput

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.Println("successfully signed in")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})
}
