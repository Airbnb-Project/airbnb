package handler

import (
	"airbnb/feature/user"
	"airbnb/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userController{
		srv: us,
	}
}

func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		cnv := ReqToCore(input)
		_, err := uc.srv.Register(*cnv)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success create account"))
	}
}

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		if input.Email == "" {
			return c.JSON(helper.ErrorResponse("email is empty"))
		} else if input.Password == "" {
			return c.JSON(helper.ErrorResponse("password is empty"))
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success login", LoginResponse(res), token))
	}
}

func (uc *userController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show profile", ProfileResponse(res)))
	}
}

func (uc *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		token := c.Get("user")

		formHeader, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(helper.ErrorResponse("please upload an image"))
		}

		cnv := ReqToCore(input)
		_, err = uc.srv.Update(token, formHeader, *cnv)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success update profile"))
	}
}

func (uc *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		err := uc.srv.Delete(token)
		if err != nil {
			c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success delete account"))
	}
}
