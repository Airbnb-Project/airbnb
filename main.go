package main

import (
	"airbnb/config"
	homeData "airbnb/feature/homestay/data"
	homeHdl "airbnb/feature/homestay/handler"
	homeSrv "airbnb/feature/homestay/service"
	usrData "airbnb/feature/user/data"
	usrHdl "airbnb/feature/user/handler"
	usrSrv "airbnb/feature/user/service"
	"airbnb/helper"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	v := validator.New()
	cld := helper.NewCloudinary(cfg)
	// coreapiMidtrans := helper.NewCoreMidtrans(cfg)

	config.Migrate(db)

	// SETUP
	uData := usrData.New(db)
	uSrv := usrSrv.New(uData, v, cld)
	uHdl := usrHdl.New(uSrv)

	hData := homeData.New(db)
	hSrv := homeSrv.New(hData, v, cld)
	hHdl := homeHdl.New(hSrv)

	// MIDDLEWARE
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	// ROUTE
	e.POST("/register", uHdl.Register())
	e.POST("/login", uHdl.Login())
	e.POST("/users", uHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/users", uHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/users", uHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/homestays", hHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/homestays", hHdl.List())
	e.POST("/homestays", hHdl.GetbyID())
	e.POST("/homestays", hHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/homestays", hHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/homestays", hHdl.Myhome(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
