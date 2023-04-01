package service

import (
	"airbnb/feature/homestay"
	"airbnb/helper"
	"errors"
	"log"
	"math"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
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
		var msg string
		if strings.Contains(err.Error(), "access denied") {
			msg = "access denied, unauthorized"
		} else {
			msg = "internal server error"
		}
		return homestay.Core{}, errors.New(msg)
	}

	return res, nil
}

func (hs *homeService) List(page int) (map[string]interface{}, []homestay.Core, error) {
	if page < 1 {
		page = 1
	}

	// limit optional, if limit change, offset change
	limit := 4
	offset := (page - 1) * limit

	totalRecord, res, err := hs.qry.List(limit, offset)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "homestay not found"
		} else {
			msg = "internal server error"
		}
		return nil, []homestay.Core{}, errors.New(msg)
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(limit)))
	if page > totalPage {
		return nil, nil, errors.New("page not found")
	}

	pagination := make(map[string]interface{})
	pagination["page"] = page
	pagination["limit"] = limit
	pagination["offset"] = offset
	pagination["totalRecord"] = totalRecord
	pagination["totalPage"] = totalPage

	return pagination, res, nil
}

func (hs *homeService) GetbyID(homestayID uint) (homestay.Core, error) {
	res, err := hs.qry.GetbyID(homestayID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "homestay not found"
		} else {
			msg = "internal server error"
		}
		return homestay.Core{}, errors.New(msg)
	}

	return res, nil
}

func (hs *homeService) Update(token interface{}, homestayID uint, updateHomestay homestay.Core, images []*multipart.FileHeader) (homestay.Core, error) {
	id := helper.ExtractToken(token)

	// check if images are exist
	if len(images) != 0 {
		// check format file/images
		for _, v := range images {
			fileimg := strings.Split(v.Filename, ".")
			format := fileimg[len(fileimg)-1]
			if format != "png" && format != "jpg" && format != "jpeg" {
				return homestay.Core{}, errors.New("file format not png, jpg, or jpeg")
			}
		}

		// upload multiple image
		image := []homestay.Image{}
		for _, v := range images {
			imgURL, err := hs.cld.Upload(v)
			if err != nil {
				log.Println("error upload image", err)
				return homestay.Core{}, errors.New("failed to upload images")
			}
			image = append(image, homestay.Image{ImageURL: imgURL})
		}
	}

	res, err := hs.qry.Update(id, homestayID, updateHomestay)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "homestay not found"
		} else if strings.Contains(err.Error(), "someone") {
			msg = "access is denied, unauthorized request"
		} else {
			msg = "internal server error"
		}
		return homestay.Core{}, errors.New(msg)
	}

	return res, nil
}

func (hs *homeService) Delete(token interface{}, homestayID uint) error {
	id := helper.ExtractToken(token)

	err := hs.qry.Delete(id, homestayID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "homestay not found"
		} else if strings.Contains(err.Error(), "someone") {
			msg = "access is denied, unauthorized request"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	return nil
}

func (hs *homeService) Myhome(token interface{}) ([]homestay.Core, error) {
	id := helper.ExtractToken(token)

	res, err := hs.qry.Myhome(id)
	if err != nil {
		log.Println(err)
		return []homestay.Core{}, errors.New("interval server error")
	}

	return res, nil
}
