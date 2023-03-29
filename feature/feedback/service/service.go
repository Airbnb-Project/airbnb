package service

import (
	"airbnb/feature/feedback"
	"airbnb/helper"
	"errors"
	"log"
	"strings"
)

type feedbackService struct {
	qry feedback.FeedbackData
}

func New(fd feedback.FeedbackData) feedback.FeedbackService {
	return &feedbackService{qry: fd}
}

func (fs *feedbackService) Add(token interface{}, homestayID uint, newFeedback feedback.Core) error {
	id := helper.ExtractToken(token)

	err := fs.qry.Add(id, homestayID, newFeedback)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil
}

func (fs *feedbackService) List(homestayID uint) ([]feedback.Core, error) {
	res, err := fs.qry.List(homestayID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "feedback not found"
		} else {
			msg = "internal server error"
		}
		return []feedback.Core{}, errors.New(msg)
	}

	return res, nil
}

func (fs *feedbackService) MyFeedback(token interface{}) ([]feedback.Core, error) {
	id := helper.ExtractToken(token)
	res, err := fs.qry.MyFeedback(id)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "feedback not found"
		} else {
			msg = "internal server error"
		}
		return []feedback.Core{}, errors.New(msg)
	}

	return res, nil
}
