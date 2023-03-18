package handler

import (
	"airbnb/feature/homestay"
	"airbnb/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type homeHandler struct {
	srv homestay.HomeService
}

func New(h homestay.HomeService) homestay.HomeHandler {
	return &homeHandler{srv: h}
}

func (hh *homeHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		// upload multiple images
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		// check if images are exist
		imageData, found := form.File["images"]
		if !found {
			return c.JSON(helper.ErrorResponse("please upload images"))
		}

		newHome := homestay.Core{}
		err = copier.Copy(&newHome, &input)
		if err != nil {
			log.Println("handler add homestay error", err.Error())
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		_, err = hh.srv.Add(token, newHome, imageData)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success add new homestay"))
	}
}

func (hh *homeHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		str := c.QueryParam("page")
		page, err := strconv.Atoi(str)
		if err != nil {
			log.Println("query param error (handler)", err.Error())
			return c.JSON(helper.ErrorResponse("page not found"))
		}

		paginate, res, err := hh.srv.List(page)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := []HomeResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler list homestay error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		// for i := range res {
		// 	resp[i].Image = res[i].Images[0].ImageURL
		// }

		pagination := helper.PaginationResponse{
			Page:        paginate["page"].(int),
			Limit:       paginate["limit"].(int),
			Offset:      paginate["offset"].(int),
			TotalRecord: paginate["totalRecord"].(int),
			TotalPage:   paginate["totalPage"].(int),
		}

		response := helper.WithPagination{
			Pagination: pagination,
			Data:       resp,
			Message:    "success show list homestay",
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (hh *homeHandler) GetbyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		homestayID, err := strconv.Atoi(param)
		if err != nil {
			log.Println("handler param get detail error", err.Error())
			return c.JSON(helper.ErrorResponse("convert id error"))
		}

		res, err := hh.srv.GetbyID(uint(homestayID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := HomeDetailResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler get detail homestay error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "succes show detail homestay", resp))
	}
}

func (hh *homeHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		param := c.Param("id")
		homestayID, err := strconv.Atoi(param)
		if err != nil {
			log.Println("handler param get detail error", err.Error())
			return c.JSON(helper.ErrorResponse("convert id error"))
		}

		input := AddRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		// upload multiple images
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		updateHome := homestay.Core{}
		err = copier.Copy(&updateHome, &input)
		if err != nil {
			log.Println("handler update homestay error", err.Error())
			return c.JSON(helper.ErrorResponse("bad request"))
		}

		_, err = hh.srv.Update(token, uint(homestayID), updateHome, form)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success update homestay"))
	}
}

func (hh *homeHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		param := c.Param("id")
		homestayID, err := strconv.Atoi(param)
		if err != nil {
			log.Println("handler param get detail error", err.Error())
			return c.JSON(helper.ErrorResponse("convert id error"))
		}

		err = hh.srv.Delete(token, uint(homestayID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success delete homestay"))
	}
}

func (hh *homeHandler) Myhome() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := hh.srv.Myhome(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		resp := []HomeResponse{}
		err = copier.Copy(&resp, &res)
		if err != nil {
			log.Println("handler list my homestay error", err.Error())
			return c.JSON(helper.ErrorResponse("failed to marshal response"))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show my homestay"))
	}
}
