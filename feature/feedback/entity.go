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
	MyFeedback() echo.HandlerFunc
}

type FeedbackService interface {
	Add(token interface{}, homestayID uint, newFeedback Core) error
	List(homestayID uint) ([]Core, error)
	MyFeedback(token interface{}) ([]Core, error)
}

type FeedbackData interface {
	Add(userID uint, homestayID uint, newFeedback Core) error
	List(homestayID uint) ([]Core, error)
	MyFeedback(userID uint) ([]Core, error)
}
