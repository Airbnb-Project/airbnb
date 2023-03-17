package handler

import (
	"airbnb/feature/homestay"
	"airbnb/helper"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type homeHandler struct {
	srv homestay.HomeService
}

func New(h homestay.HomeService) homestay.HomeHandler {
	return &homeHandler{srv: h}
}

func (hh *homeHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		// upload multiple images
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		// check if images are exist
		imageData, found := form.File["images"]
		if !found {
			return c.JSON(helper.ErrorResponse("please upload images"))
		}

		newHome := homestay.Core{}
		err = copier.Copy(&newHome, &input)
		if err != nil {
			log.Println("handler add homestay error", err)
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		_, err = hh.srv.Add(token, newHome, imageData)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success add new homestay"))
	}
}

func (hh *homeHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (hh *homeHandler) GetbyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (hh *homeHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (hh *homeHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (hh *homeHandler) Myhome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
