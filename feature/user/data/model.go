package data

import (
	"airbnb/feature/homestay/data"
	"airbnb/feature/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	Password       string
	Role           string
	ProfilePicture string
	Phone          string
	Address        string
	Homestay       []data.Homestay // `gorm:"foreignKey:UserID"`
}

func DataToCore(data User) user.Core {
	return user.Core{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		Role:           data.Role,
		ProfilePicture: data.ProfilePicture,
		Phone:          data.Phone,
		Address:        data.Address,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:          gorm.Model{ID: data.ID},
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		Role:           data.Role,
		ProfilePicture: data.ProfilePicture,
		Phone:          data.Phone,
		Address:        data.Address,
	}
}
