package data

import (
	"airbnb/feature/homestay"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	UserID    uint
	Name      string
	Address   string
	Phone     string
	Facility  string
	Price     int
	Images    []Image    // `gorm:"foreignKey:HomestayID"`
	Feedbacks []Feedback // `gorm:"foreignKey:HomestayID"`
}

type Image struct {
	gorm.Model
	HomestayID uint
	ImageURL   string
}

type Feedback struct {
	gorm.Model
	Rating uint
	Note   string
}

func DataToCore(data Homestay) homestay.Core {
	img := []homestay.Image{}
	for _, v := range data.Images {
		img = append(img, homestay.Image{
			ID:       v.ID,
			ImageURL: v.ImageURL,
		})
	}

	return homestay.Core{
		ID:        data.ID,
		Name:      data.Name,
		Address:   data.Address,
		Phone:     data.Phone,
		Facility:  data.Facility,
		Price:     data.Price,
		Images:    img,
		Feedbacks: []homestay.Feedback{},
	}
}

func CoreToData(data homestay.Core) Homestay {
	img := []Image{}
	for _, v := range data.Images {
		img = append(img, Image{
			Model:      gorm.Model{ID: data.ID},
			HomestayID: v.ID,
			ImageURL:   v.ImageURL,
		})
	}

	return Homestay{
		Model:     gorm.Model{ID: data.ID},
		UserID:    data.UserID,
		Name:      data.Name,
		Address:   data.Address,
		Phone:     data.Phone,
		Facility:  data.Facility,
		Price:     data.Price,
		Images:    img,
		Feedbacks: []Feedback{},
	}
}
