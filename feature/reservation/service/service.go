package service

import (
	"airbnb/feature/reservation"
	"airbnb/helper"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	qry reservation.ReservationData
	vld *validator.Validate
	pay helper.PaymentGateway
}

func New(rd reservation.ReservationData, v *validator.Validate, p helper.PaymentGateway) reservation.ReservationService {
	return &reservationService{
		qry: rd,
		vld: v,
		pay: p,
	}
}

func (rs *reservationService) Create(token interface{}, newReservation reservation.Core) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rs *reservationService) List(token interface{}) ([]reservation.Core, error) {
	return []reservation.Core{}, nil
}

func (rs *reservationService) Detail(token interface{}, reservationID uint) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rs *reservationService) Accept(token interface{}, reservationID uint, status string) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rs *reservationService) Cancel(token interface{}, reservationID uint, status string) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rs *reservationService) Callback(ticket string, status string) error {
	return nil
}
