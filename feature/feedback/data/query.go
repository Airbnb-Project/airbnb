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

func (fq *feedbackQuery) Add(userID uint, homestayID uint, newFeedback feedback.Core) error {
	cnv := CoreToData(newFeedback)
	cnv.UserID = userID
	cnv.HomestayID = homestayID
	err := fq.db.Create(&cnv).Error
	if err != nil {
		log.Println("add feedback query error", err.Error())
		return errors.New("cannot add feedback")
	}

	newFeedback.ID = cnv.ID

	return nil
}

func (fq *feedbackQuery) List(homestayID uint) ([]feedback.Core, error) {
	fb := []Feedback{}
	err := fq.db.Raw("SELECT f.id, f.rating, f.note, u.id, u.name, h.id, h.name, h.address FROM feedbacks f JOIN users u ON u.id = f.user_id JOIN homestays h ON h.id = f.homestay_id WHERE homestay_id = ? AND deleted_at is NULL", homestayID).Error
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
	err := fq.db.Where("user_id = ?", userID).Find(&fb).Error
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
