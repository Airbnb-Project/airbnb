package data

import (
	"airbnb/feature/feedback"
	"airbnb/feature/homestay/data"
	"errors"
	"log"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackData {
	return &feedbackQuery{db: db}
}

func (fq *feedbackQuery) Add(userID uint, homestayID uint, newFeedback feedback.Core) error {
	// find homestay data for insert feedback
	hs := data.Homestay{}
	err := fq.db.Where("id = ?", homestayID).First(&hs).Error
	if err != nil {
		log.Println("feedback find homestay data query error", err.Error())
		return errors.New("cannot find homestay data")
	}

	cnv := CoreToData(newFeedback)
	cnv.UserID = userID
	cnv.HomestayID = homestayID
	err = fq.db.Create(&cnv).Error
	if err != nil {
		log.Println("add feedback query error", err.Error())
		return errors.New("cannot add feedback")
	}

	newFeedback.ID = cnv.ID

	return nil
}

func (fq *feedbackQuery) List() ([]feedback.Core, error) {
	fb := []Feedback{}
	err := fq.db.Order("created_at DESC").Find(&fb).Error
	if err != nil {
		log.Println("show feedback query error", err.Error())
		return []feedback.Core{}, errors.New("data not found, cannot show list feedback")
	}

	list := []feedback.Core{}
	for _, v := range fb {
		list = append(list, DataToCore(v))
	}

	return list, nil
}

func (fq *feedbackQuery) MyFeedback(userID uint) ([]feedback.Core, error) {
	fb := []Feedback{}
	err := fq.db.Where("id = ?", userID).Find(&fb).Error
	if err != nil {
		log.Println("show my feedback query error", err.Error())
		return []feedback.Core{}, errors.New("data not found, cannot find my feedback")
	}

	list := []feedback.Core{}
	for _, v := range fb {
		list = append(list, DataToCore(v))
	}

	return list, nil
}
