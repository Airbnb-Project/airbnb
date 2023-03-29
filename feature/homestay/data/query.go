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

	var role string
	err := hd.db.Raw("SELECT role FROM users WHERE user_id = ?", userID).First(&role).Error
	if err != nil {
		log.Println("query role homestay error", err.Error())
		return homestay.Core{}, errors.New("access denied")
	}

	if role != "host" {
		return homestay.Core{}, errors.New("access denied")
	}

	err = hd.db.Create(&cnv).Error
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
	var totalRecord int64

	err := hd.db.Table("Homestay").Count(&totalRecord).Error
	if err != nil {
		log.Println("find totalrecord query error", err.Error())
		return 0, []homestay.Core{}, errors.New("data not found")
	}

	err = hd.db.Limit(limit).Offset(offset).Order("id DESC").Find(&hs).Error
	if err != nil {
		log.Println("show list query error", err.Error())
		return 0, []homestay.Core{}, errors.New("data not found, cannot show list homestay")
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

	return int(totalRecord), list, nil
}

func (hd *homeData) GetbyID(homestayID uint) (homestay.Core, error) {
	hs := Homestay{}
	err := hd.db.Where("id = ?", homestayID).Preload("Feedback").First(&hs).Error
	if err != nil {
		log.Println("show by id query error", err.Error())
		return homestay.Core{}, errors.New("data not found, cannot show detail homestay")
	}

	return DataToCore(hs), nil
}

func (hd *homeData) Update(userID uint, homestayID uint, updateHomestay homestay.Core) (homestay.Core, error) {
	// check if this is the owner wants to edit the homestay
	hs := Homestay{}
	err := hd.db.Where("user_id = ? and id = ?", userID, homestayID).First(&hs).Error
	if err != nil {
		log.Println("check owner query error", err.Error())
		return homestay.Core{}, errors.New("cannot edit someone homestay")
	}

	cnv := CoreToData(updateHomestay)
	qry := hd.db.Where("user_id = ? and id = ?", userID, homestayID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return homestay.Core{}, errors.New("no updated data")
	}

	err = qry.Error
	if err != nil {
		log.Println("update homestay query error", err.Error())
		return homestay.Core{}, errors.New("data not found, cannot update homestay")
	}

	return updateHomestay, nil
}

func (hd *homeData) Delete(userID uint, homestayID uint) error {
	// check if this is the owner wants to delete the homestay
	hs := Homestay{}
	err := hd.db.Where("user id = ?", userID).First(&hs).Error
	if err != nil {
		log.Println("check owner query error", err.Error())
		return errors.New("cannot delete someone homestay")
	}

	qry := hd.db.Delete(&hs, homestayID)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no updated data")
	}

	err = qry.Error
	if err != nil {
		log.Println("delete homestay query error", err.Error())
		return errors.New("data not found, cannot delete homestay")
	}

	return nil
}

func (hd *homeData) Myhome(userID uint) ([]homestay.Core, error) {
	hs := []Homestay{}
	err := hd.db.Where("id = ?", userID).Find(&hs).Error
	if err != nil {
		log.Println("show myhomestay query error", err.Error())
		return []homestay.Core{}, errors.New("cannot show my list homestay")
	}

	list := []homestay.Core{}
	for _, v := range hs {
		list = append(list, DataToCore(v))
	}

	return list, nil
}
