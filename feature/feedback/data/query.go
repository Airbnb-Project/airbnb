package data

import (
	"airbnb/feature/feedback"
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

func (fq *feedbackQuery) Add(userID uint, newFeedback feedback.Core) error {
	cnv := CoreToData(newFeedback)
	cnv.UserID = userID
	err := fq.db.Create(&cnv).Error
	if err != nil {
		log.Println("add feedback query error", err.Error())
		return errors.New("cannot add feedback")
	}

	newFeedback.ID = cnv.ID

	return nil
}

func (fq *feedbackQuery) List(userID uint, homestayID uint) ([]feedback.Core, error) {
	fb := []Feedback{}
	err := fq.db.Where("homestay_id = ?", homestayID).Order("created_at DESC").Find(&fb).Error
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
