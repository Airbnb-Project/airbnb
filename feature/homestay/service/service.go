package service

import (
	"airbnb/feature/homestay"
	"airbnb/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"

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
	id := helper.ExtractToken(token)

	err := hs.vld.Struct(newHomestay)
	if err != nil {
		log.Println("error", err)
		msg := helper.ValidationErrorHandle(err)
		return homestay.Core{}, errors.New(msg)
	}

	// check format file/images
	for _, v := range imagesData {
		fileimg := strings.Split(v.Filename, ".")
		format := fileimg[len(fileimg)-1]
		if format != "png" && format != "jpg" && format != "jpeg" {
			return homestay.Core{}, errors.New("file format not png, jpg, or jpeg")
		}
	}

	// upload multiple image
	image := []homestay.Image{}
	for _, v := range imagesData {
		imgURL, err := hs.cld.Upload(v)
		if err != nil {
			log.Println("error upload image", err)
			return homestay.Core{}, errors.New("failed to upload images")
		}
		image = append(image, homestay.Image{ImageURL: imgURL})
	}

	newHomestay.Images = image

	res, err := hs.qry.Add(id, newHomestay)
	if err != nil {
		{
			return homestay.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
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
