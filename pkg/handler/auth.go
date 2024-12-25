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

// @Summary		SignUp
// @Tags			Auth
// @Description	Create account
// @Accept			json
// @Produce		json
// @Param			input	body		musiclib.User	true	"Account info"
// @Success		200		{object}	RegisterResponce
// @Failure		400,404	{object}	MyError	"4** error"
// @Failure		500		{object}	MyError	"5** error"
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var input musiclib.User

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	if err := validate.Struct(input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	logrus.Printf("successfully signed up, userId=%v", id)
	return ctx.Status(http.StatusOK).JSON(RegisterResponce{
		ID:          id,
		AccessToken: token,
	})
}

type SignInResponce struct {
	AccessToken string `json:"access_token"`
}

// @Summary		SignIn
// @Tags			Auth
// @Description	Login
// @Accept			json
// @Produce		json
// @Param			input	body		SignInInput		true	"Credentials"
// @Success		200		{object}	SignInResponce	"Successful login"
// @Failure		400,404	{object}	MyError			"4** error"
// @Failure		500		{object}	MyError			"5** error"
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(ctx *fiber.Ctx) error {
	var input SignInInput

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	if err := validate.Struct(input); err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(MyError{
			Err: err.Error(),
		})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Error(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(MyError{
			Err: err.Error(),
		})
	}

	logrus.Println("successfully signed in")
	return ctx.Status(http.StatusOK).JSON(SignInResponce{
		AccessToken: token,
	})
}
