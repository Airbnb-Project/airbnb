package data

import (
	"airbnb/feature/reservation"

	"gorm.io/gorm"
)

type reservationData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationData {
	return &reservationData{db: db}
}

func (rd *reservationData) Create(userID uint, newReservation reservation.Core) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rd *reservationData) List(userID uint) ([]reservation.Core, error) {
	return []reservation.Core{}, nil
}

func (rd *reservationData) Detail(userID uint, reservationID uint) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rd *reservationData) Update(userID uint, reservationID uint, status string) (reservation.Core, error) {
	return reservation.Core{}, nil
}

func (rd *reservationData) Callback(tiket string, status string) error {
	return nil
}
