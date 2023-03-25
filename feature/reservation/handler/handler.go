package handler

import (
	"airbnb/feature/reservation"

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
		return nil
	}
}

func (rh *reservationHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (rh *reservationHandler) Detail() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
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
