package data

import (
	"airbnb/feature/reservation"
	"errors"
	"log"

	"gorm.io/gorm"
)

type reservationData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationData {
	return &reservationData{db: db}
}

func (rd *reservationData) Create(userID uint, newReservation reservation.Core) (reservation.Core, error) {
	cnv := CoreToData(newReservation)
	err := rd.db.Create(&cnv).Error
	if err != nil {
		log.Println("query create reservation error", err.Error())
		return reservation.Core{}, errors.New("cannot create reservation")
	}

	newReservation.ID = cnv.ID

	return newReservation, nil
}

func (rd *reservationData) List(userID uint) ([]reservation.Core, error) {
	rsv := []ReservationModel{}
	err := rd.db.Where("user_id = ?", userID).Find(&rsv).Error
	if err != nil {
		log.Println("query list reservation error", err.Error())
		return []reservation.Core{}, errors.New("data not found, cannot show list reservation")
	}

	list := []reservation.Core{}
	for _, v := range rsv {
		list = append(list, ToCore(v))
	}

	return list, nil
}

func (rd *reservationData) Detail(userID uint, reservationID uint) (reservation.Core, error) {
	rsv := ReservationModel{}
	err := rd.db.Where("user_id = ? and id = ?", userID, reservationID).First(&rsv).Error
	if err != nil {
		log.Println("query detail reservation error", err.Error())
		return reservation.Core{}, errors.New("data not found, cannot show detail reservation")
	}

	return ToCore(rsv), nil
}

func (rd *reservationData) Update(userID uint, reservationID uint, status string) (reservation.Core, error) {
	err := rd.db.Raw("SELECT role FROM user WHERE user_id = ?", userID).Error
	if err != nil {
		log.Println("query role reservation update error", err.Error())
		return reservation.Core{}, errors.New("access denied")
	}

	rsv := ReservationModel{}
	qry := rd.db.Model(&rsv).Where("user_id = ? and id = ?", userID, reservationID).Update("status", status)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return reservation.Core{}, errors.New("no updated data")
	}

	err = qry.Error
	if err != nil {
		log.Println("query update reservation error", err.Error())
		return reservation.Core{}, errors.New("cannot update reseration")
	}

	return ToCore(rsv), nil
}

func (rd *reservationData) Callback(ticket string, status string) error {
	rsv := Reservation{}
	qry := rd.db.Model(&rsv).Where("ticket = ?", ticket).Update("status", status)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no updated data")
	}

	err := qry.Error
	if err != nil {
		log.Println("query callback error", err.Error())
		return errors.New("failed to callback")
	}

	return nil
}
