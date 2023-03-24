package data

import (
	"airbnb/feature/reservation"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID          uint
	HomestayID      uint
	Ticket          string
	Checkin         string
	Checkout        string
	ReservationDate string
	TotalPrice      int
	Status          string
	PaymentLink     string
	Bank            string
	VAnumber        string
}

func DataToCore(data Reservation) reservation.Core {
	return reservation.Core{
		ID:              data.ID,
		UserID:          data.UserID,
		HomestayID:      data.HomestayID,
		Ticket:          data.Ticket,
		Checkin:         data.Checkin,
		Checkout:        data.Checkout,
		ReservationDate: data.ReservationDate,
		TotalPrice:      data.TotalPrice,
		Status:          data.Status,
		PaymentLink:     data.PaymentLink,
		Bank:            data.Bank,
		VAnumber:        data.VAnumber,
	}
}

func CoreToData(data reservation.Core) Reservation {
	return Reservation{
		Model:           gorm.Model{ID: data.ID},
		UserID:          data.UserID,
		HomestayID:      data.HomestayID,
		Ticket:          data.Ticket,
		Checkin:         data.Checkin,
		Checkout:        data.Checkout,
		ReservationDate: data.ReservationDate,
		TotalPrice:      data.TotalPrice,
		Status:          data.Status,
		PaymentLink:     data.PaymentLink,
		Bank:            data.Bank,
		VAnumber:        data.VAnumber,
	}
}
