package handler

import (
	"airbnb/feature/reservation"
	"airbnb/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type reservationHandler struct {
	srv reservation.ReservationService
}

func New(r reservation.ReservationService) reservation.ReservationHandler {
	return &reservationHandler{srv: r}
}

func (rh *reservationHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := rsvRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		newReservation := reservation.Core{}
		err := copier.Copy(&newReservation, &input)
		if err != nil {
			log.Println("handler create reservation error", err.Error())
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		res, err := rh.srv.Create(token, newReservation)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success create reservation", res.ID))
	}
}

func (rh *reservationHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := rh.srv.List(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := []RsvResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler list reservation error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		for _, v := range res {
			resp = append(resp, RsvResponse{
				ID:              v.ID,
				Ticket:          v.Ticket,
				Name:            v.Name,
				Image:           v.Image,
				TotalPrice:      v.TotalPrice,
				ReservationDate: v.ReservationDate,
			})
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show list reservation", resp))
	}
}

func (rh *reservationHandler) Detail() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		reservationID, err := strconv.Atoi(param)
		if err != nil {
			log.Println("handler param get detail error", err.Error())
			return c.JSON(helper.ErrorResponse("convert id error"))
		}

		token := c.Get("user")
		res, err := rh.srv.Detail(token, uint(reservationID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := RsvDetailResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler get detail reservation error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show detail reservation", resp))
	}
}

func (rh *reservationHandler) Accept() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (rh *reservationHandler) Cancel() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (rh *reservationHandler) Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
