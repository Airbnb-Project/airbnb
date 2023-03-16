package data

import (
	"airbnb/feature/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	// check if user is duplicated
	dupUser := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupUser).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email already registered")
	}

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query register error", err.Error())
		return user.Core{}, errors.New("cannot create new user")
	}
	// give id to newuser
	newUser.ID = cnv.ID

	return newUser, nil
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	// check if email is empty
	if email == "" {
		log.Println("empty email")
		return user.Core{}, errors.New("email is empty")
	}

	usr := User{}
	err := uq.db.Where("email = ?", email).First(&usr).Error
	if err != nil {
		log.Println("query login error", err.Error())
		return user.Core{}, errors.New("cannot login")
	}

	return DataToCore(usr), nil
}

func (uq *userQuery) Profile(userID uint) (user.Core, error) {
	usr := User{}
	err := uq.db.Where("id = ?", userID).First(&usr).Error
	if err != nil {
		log.Println("query get profile error", err.Error())
		return user.Core{}, errors.New("cannot show profile")
	}

	return DataToCore(usr), nil
}

func (uq *userQuery) Update(userID uint, updateData user.Core) (user.Core, error) {
	cnv := CoreToData(updateData)
	usr := User{}
	qry := uq.db.Model(&usr).Where("id = ?", userID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return user.Core{}, errors.New("no updated data")
	}

	err := qry.Error
	if err != nil {
		log.Println("query update error", err.Error())
		return user.Core{}, errors.New("cannot update profile")
	}

	return DataToCore(cnv), nil
}

func (uq *userQuery) Delete(userID uint) error {
	usr := User{}
	qry := uq.db.Delete(&usr, userID)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no updated data")
	}

	err := qry.Error
	if err != nil {
		log.Println("query delete error", err.Error())
		return errors.New("cannot delete account")
	}

	return nil
}
