package handler

import (
	"airbnb/feature/user"
	"mime/multipart"
)

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	Phone          string `json:"phone" form:"phone"`
	Address        string `json:"address" form:"address"`
	Role           string `json:"role" form:"role"`
	ProfilePicture string `json:"photo" form:"photo"`
	FileHeader     multipart.FileHeader
}

func ReqToCore(data interface{}) *user.Core {
	core := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		core.Name = cnv.Name
		core.Email = cnv.Email
		core.Password = cnv.Password
		core.Phone = cnv.Phone
		core.Address = cnv.Address
		core.Role = cnv.Role
	case LoginRequest:
		cnv := data.(LoginRequest)
		core.Email = cnv.Email
		core.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		core.Name = cnv.Name
		core.Email = cnv.Email
		core.Password = cnv.Password
		core.Phone = cnv.Phone
		core.Address = cnv.Address
		core.Role = cnv.Role
		core.ProfilePicture = cnv.ProfilePicture
	default:
		return nil
	}

	return &core
}
