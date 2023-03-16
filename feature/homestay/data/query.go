package data

import (
	"airbnb/feature/homestay"

	"gorm.io/gorm"
)

type homeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.HomeData {
	return &homeData{db: db}
}

func (hd *homeData) Add(userID uint, newHomestay homestay.Core) (homestay.Core, error) {
	cnv := CoreToData(newHomestay)
	cnv.UserID = userID
	hd.db.Create(&cnv)
	return homestay.Core{}, nil
}

func (hd *homeData) List(limit int, offset int) ([]homestay.Core, error) {
	return []homestay.Core{}, nil
}

func (hd *homeData) GetbyID(homestayID uint) (homestay.Core, error) {
	return homestay.Core{}, nil
}

func (hd *homeData) Update(userID uint, homestayID uint, updateHomestay homestay.Core) (homestay.Core, error) {
	return homestay.Core{}, nil
}

func (hd *homeData) Delete(userID uint, homestayID uint) error {
	return nil
}

func (hd *homeData) Myhome(userID uint) ([]homestay.Core, error) {
	return []homestay.Core{}, nil
}
