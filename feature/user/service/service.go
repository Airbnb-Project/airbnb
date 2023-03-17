package service

import (
	"airbnb/feature/user"
	"airbnb/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
	cld helper.Uploader
}

func New(ud user.UserData, v *validator.Validate, u helper.Uploader) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: v,
		cld: u,
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	hashed, err := helper.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error", err.Error())
		return user.Core{}, errors.New("password process error")
	}

	// validation for register
	err = uuc.vld.Struct(&newUser)
	if err != nil {
		log.Println("error", err)
		msg := helper.ValidationErrorHandle(err)
		return user.Core{}, errors.New(msg)
	}

	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		var msg string
		if strings.Contains(err.Error(), "registered") {
			msg = "data already used"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)

	}

	return res, nil
}

func (uuc *userUseCase) Login(email string, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(email)
	if err != nil {
		var msg string
		if strings.Contains(err.Error(), "empty") {
			msg = "email or password not allowed empty"
		} else {
			msg = "account not registered or server error"
		}
		return "", user.Core{}, errors.New(msg)
	}

	err = helper.CheckPassword(res.Password, password)
	if err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}

	useToken, _ := helper.GenerateToken(int(res.ID))
	return useToken, res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	res, err := uuc.qry.Profile(id)
	if err != nil {
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Update(token interface{}, fileData *multipart.FileHeader, updateData user.Core) (user.Core, error) {
	id := helper.ExtractToken(token)

	if updateData.Password != "" {
		hashed, _ := helper.GeneratePassword(updateData.Password)
		updateData.Password = string(hashed)
	}

	// check if image already exist
	if fileData != nil {
		secureURL, err := uuc.cld.Upload(fileData)
		if err != nil {
			log.Println("error upload file", err)
			var msg string
			if strings.Contains(err.Error(), "bad request") {
				msg = "bad request from user"
			} else {
				msg = "failed to upload image, internal server error"
			}
			return user.Core{}, errors.New(msg)
		}
		updateData.ProfilePicture = secureURL
	}

	res, err := uuc.qry.Update(id, updateData)
	if err != nil {
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	err := uuc.qry.Delete(id)
	if err != nil {
		log.Println("delete error", err.Error())
		return errors.New("failed to delete account")
	}

	return nil
}
