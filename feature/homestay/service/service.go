package service

import (
	"airbnb/feature/homestay"
	"airbnb/helper"
	"mime/multipart"

	"github.com/go-playground/validator"
)

type homeService struct {
	qry homestay.HomeData
	vld *validator.Validate
	cld helper.Uploader
}

func New(hd homestay.HomeData, v *validator.Validate, u helper.Uploader) homestay.HomeService {
	return &homeService{
		qry: hd,
		vld: v,
		cld: u,
	}
}

func (hs *homeService) Add(token interface{}, newHomestay homestay.Core, imagesData []*multipart.FileHeader) (homestay.Core, error) {
	return homestay.Core{}, nil
}

func (hs *homeService) List(page int) (map[string]interface{}, []homestay.Core, error) {
	return make(map[string]interface{}), []homestay.Core{}, nil
}

func (hs *homeService) GetbyID(homestayID uint) (homestay.Core, error) {
	return homestay.Core{}, nil
}

func (hs *homeService) Update(token interface{}, homestayID uint, updateHomestay homestay.Core, images []*multipart.FileHeader) (homestay.Core, error) {
	return homestay.Core{}, nil
}

func (hs *homeService) Delete(token interface{}, homestayID uint) error {
	return nil
}

func (hs *homeService) Myhome(token interface{}) ([]homestay.Core, error) {
	return []homestay.Core{}, nil
}
