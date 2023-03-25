package reservation

import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint
	UserID          uint
	HomestayID      uint
	Ticket          string
	Checkin         string `validate:"required"`
	Checkout        string `validate:"required"`
	ReservationDate string
	TotalPrice      int `validate:"required"`
	Status          string
	PaymentLink     string
	Bank            string `validate:"required"`
	VAnumber        string
}

type ReservationHandler interface {
	Create() echo.HandlerFunc
	List() echo.HandlerFunc
	Detail() echo.HandlerFunc
	Accept() echo.HandlerFunc
	Cancel() echo.HandlerFunc
	Callback() echo.HandlerFunc
}

type ReservationService interface {
	Create(token interface{}, newReservation Core) (Core, error)
	List(token interface{}) ([]Core, error)
	Detail(token interface{}, reservationID uint) (Core, error)
	Accept(token interface{}, reservationID uint, status string) (Core, error)
	Cancel(token interface{}, reservationID uint, status string) (Core, error)
	Callback(ticket string, status string) error
}

type ReservationData interface {
	Create(userID uint, newReservation Core) (Core, error)
	List(userID uint) ([]Core, error)
	Detail(userID uint, reservationID uint) (Core, error)
	Update(userID uint, reservationID uint, status string) (Core, error)
	Callback(ticket string, status string) error
}
