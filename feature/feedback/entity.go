package feedback

import "github.com/labstack/echo/v4"

type Core struct {
	ID         uint
	Rating     uint
	Note       string
	UserID     uint
	HomestayID uint
}

type FeedbackHandler interface {
	Add() echo.HandlerFunc
	List() echo.HandlerFunc
}

type FeedbackService interface {
	Add(token interface{}, homestayID uint, newFeedback Core) error
	List(token interface{}, homestayID uint) ([]Core, error)
}

type FeedbackData interface {
	Add(userID uint, homestayID uint, newFeedback Core) error
	List(userID uint, homestayID uint) ([]Core, error)
}
