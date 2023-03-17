package data

import (
	"airbnb/feature/homestay"
	"errors"
	"log"

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
	// hd.db.Create(&cnv)
	err := hd.db.Create(&cnv).Error
	if err != nil {
		log.Println("query add homestay error", err.Error())
		return homestay.Core{}, errors.New("cannot create homestay")
	}

	for _, v := range cnv.Images {
		err = hd.db.Exec("INSERT INTO images(homestay_id, image_url) VALUES(?, ?)", v.HomestayID, v.ImageURL).Error
		if err != nil {
			log.Println("query add image error", err.Error())
			return homestay.Core{}, errors.New("cannot insert image")
		}

	}

	newHomestay.ID = cnv.ID

	return newHomestay, nil
}

func (hd *homeData) List(limit int, offset int) (int, []homestay.Core, error) {
	hs := []Homestay{}
	err := hd.db.Limit(limit).Offset(offset).Order("created_at DESC, id DESC").Find(&hs).Error
	if err != nil {
		log.Println("show list query error", err.Error())
		return 0, []homestay.Core{}, errors.New("cannot show list homestay")
	}

	// find homestay images
	img := Image{}
	for _, v := range hs {
		err := hd.db.Where("id = ?", v.ID).First(&img).Error
		if err != nil {
			log.Println("find image home query error", err.Error())
			return 0, []homestay.Core{}, errors.New("cannot find image from homestay")
		}
	}

	list := []homestay.Core{}
	for _, v := range hs {
		list = append(list, DataToCore(v))
	}
	var totalRecord int

	return totalRecord, list, nil
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
