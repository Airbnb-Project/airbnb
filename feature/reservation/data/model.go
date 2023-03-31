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
	Bank            string
	VAnumber        string
}

type ReservationModel struct {
	ID              uint
	UserID          uint
	HomestayID      uint
	Name            string
	Image           string
	Address         string
	Ticket          string
	Checkin         string
	Checkout        string
	Guest           int
	ReservationDate string
	TotalPrice      int
	Status          string
	Bank            string
	VAnumber        string
}

func ToCore(data ReservationModel) reservation.Core {
	return reservation.Core{
		ID:              data.ID,
		UserID:          data.UserID,
		HomestayID:      data.HomestayID,
		Name:            data.Name,
		Image:           data.Image,
		Address:         data.Address,
		Ticket:          data.Ticket,
		Checkin:         data.Checkin,
		Checkout:        data.Checkout,
		Guest:           data.Guest,
		ReservationDate: data.ReservationDate,
		TotalPrice:      data.TotalPrice,
		Status:          data.Status,
		Bank:            data.Bank,
		VAnumber:        data.VAnumber,
	}
}

func ListToCore(data []ReservationModel) []reservation.Core {
	core := []reservation.Core{}

	for _, v := range data {
		core = append(core, ToCore(v))
	}

	return core
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
		Bank:            data.Bank,
		VAnumber:        data.VAnumber,
	}
}
