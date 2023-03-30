package handler

import (
	"airbnb/feature/feedback"
	"airbnb/helper"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	srv feedback.FeedbackService
}

func New(f feedback.FeedbackService) feedback.FeedbackHandler {
	return &feedbackHandler{srv: f}
}

func (fh *feedbackHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		newFeedback := feedback.Core{}
		err := copier.Copy(&newFeedback, &input)
		if err != nil {
			log.Println("handler add feedback error", err.Error())
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		err = fh.srv.Add(token, input.HomestayID, newFeedback)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success add feedback"))
	}
}

func (fh *feedbackHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ListRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		fb := feedback.Core{}
		err := copier.Copy(&fb, &input)
		if err != nil {
			log.Println("handler list feedback error", err.Error())
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		res, err := fh.srv.List(input.HomestayID)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := []FeedbackResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler list feedback error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		for _, v := range res {
			resp = append(resp, FeedbackResponse{
				ID:              v.ID,
				Rating:          v.Rating,
				Note:            v.Note,
				UserID:          v.UserID,
				UserName:        v.User.Name,
				HomestayID:      v.HomestayID,
				HomestayName:    v.Homestay.Name,
				HomestayAddress: v.Homestay.Address,
			})
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show list feedback", resp))
	}
}

func (fh *feedbackHandler) MyFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := fh.srv.MyFeedback(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := []FeedbackResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler my feedback error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		for _, v := range res {
			resp = append(resp, FeedbackResponse{
				ID:              v.ID,
				Rating:          v.Rating,
				Note:            v.Note,
				UserID:          v.UserID,
				UserName:        v.User.Name,
				HomestayID:      v.HomestayID,
				HomestayName:    v.Homestay.Name,
				HomestayAddress: v.Homestay.Address,
			})
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show my feedback", resp))
	}
}
