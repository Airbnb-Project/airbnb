package service

import (
	"airbnb/feature/reservation"
	"airbnb/helper"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

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
	id := helper.ExtractToken(token)

	err := rs.vld.Struct(newReservation)
	if err != nil {
		log.Println("errpr", err)
		msg := helper.ValidationErrorHandle(err)
		return reservation.Core{}, errors.New(msg)
	}

	// set some default for transaction
	newReservation.Ticket = fmt.Sprintf("INV-%d-%s", id, time.Now().Format("20060102-150405"))
	newReservation.Status = "waiting for payment"
	newReservation.ReservationDate = time.Now().Format("2006-01-02")

	// charge transaction to midtrans to get the response
	vaNumber, err := rs.pay.ChargeTransaction(newReservation.Ticket, newReservation.TotalPrice, newReservation.Bank)
	if err != nil {
		log.Println(err)
		return reservation.Core{}, errors.New("charge transaction failed due to internal server error, please try again")
	}

	newReservation.VAnumber = vaNumber

	res, err := rs.qry.Create(id, newReservation)
	if err != nil {
		log.Println(err)
		return reservation.Core{}, errors.New("internal server error")
	}

	return res, nil
}

func (rs *reservationService) List(token interface{}) ([]reservation.Core, error) {
	id := helper.ExtractToken(token)
	res, err := rs.qry.List(id)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "reservation not found"
		} else {
			msg = "internal server error"
		}
		return []reservation.Core{}, errors.New(msg)
	}

	return res, nil
}

func (rs *reservationService) Detail(token interface{}, reservationID uint) (reservation.Core, error) {
	id := helper.ExtractToken(token)
	res, err := rs.qry.Detail(id, reservationID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "reservation not found"
		} else {
			msg = "internal server error"
		}
		return reservation.Core{}, errors.New(msg)
	}

	return res, nil
}

func (rs *reservationService) Accept(token interface{}, reservationID uint, status string) (reservation.Core, error) {
	id := helper.ExtractToken(token)
	var role string

	res, err := rs.qry.Update(id, reservationID, status)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "reservation not found"
		} else {
			msg = "internal server error"
		}
		return reservation.Core{}, errors.New(msg)
	}

	if role != "host" {
		return reservation.Core{}, errors.New("access denied")
	}
	return res, nil
}

func (rs *reservationService) Cancel(token interface{}, reservationID uint, status string) (reservation.Core, error) {
	id := helper.ExtractToken(token)
	res, err := rs.qry.Update(id, reservationID, status)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "reservation not found"
		} else {
			msg = "internal server error"
		}
		return reservation.Core{}, errors.New(msg)
	}

	return res, nil
}

func (rs *reservationService) Callback(ticket string, status string) error {
	err := rs.qry.Callback(ticket, status)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "callback") {
			msg = "failed to callback"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	return nil
}
