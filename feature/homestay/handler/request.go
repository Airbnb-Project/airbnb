package handler

import "mime/multipart"

type AddRequest struct {
	Name       string  `json:"name" form:"name"`
	Address    string  `json:"address" form:"address"`
	Phone      string  `json:"phone" form:"phone"`
	Price      float64 `json:"price" form:"price"`
	Facility   string  `json:"facility" form:"facility"`
	FileHeader []multipart.FileHeader
}
