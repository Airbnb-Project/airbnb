package homestay

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	UserID   uint
	Name     string `validate:"required"`
	Address  string `validate:"min=5"`
	Phone    string
	Facility string
	Images   []Image
}

type Image struct {
	ID       uint
	ImageURL string
}

type HomeHandler interface {
	Add() echo.HandlerFunc
	List() echo.HandlerFunc
	GetbyID() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Myhome() echo.HandlerFunc
}

type HomeService interface {
	Add(token interface{}, newHomestay Core, images []*multipart.FileHeader) (Core, error)
	// List(page int) (map[string]interface{}, []Core, error)
	// GetbyID(homestayID uint) (Core, error)
	// Update(token interface{}, homestayID uint, updateHomestay Core, images []*multipart.FileHeader) (Core, error)
	// Delete(token interface{}, homestayID uint) error
	// Myhome(token interface{}) ([]Core, error)
}

type HomeData interface {
	Add(userID uint, newHomestay Core) (Core, error)
	// List(limit int, offset int) ([]Core, error)
	// GetbyID(homestayID uint) (Core, error)
	// Update(userID uint, homestayID uint, updateHomestay Core) (Core, error)
	// Delete(userID uint, homestayID uint) error
	// Myhome(userID uint) ([]Core, error)
}
