package main

import (
	"airbnb/config"
	fbData "airbnb/feature/feedback/data"
	fbHdl "airbnb/feature/feedback/handler"
	fbSrv "airbnb/feature/feedback/service"
	homeData "airbnb/feature/homestay/data"
	homeHdl "airbnb/feature/homestay/handler"
	homeSrv "airbnb/feature/homestay/service"
	rsvData "airbnb/feature/reservation/data"
	rsvHdl "airbnb/feature/reservation/handler"
	rsvSrv "airbnb/feature/reservation/service"
	usrData "airbnb/feature/user/data"
	usrHdl "airbnb/feature/user/handler"
	usrSrv "airbnb/feature/user/service"
	"airbnb/helper"
	"log"

	"github.com/go-playground/validator/v10"
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
	coreapiMidtrans := helper.NewCoreMidtrans(cfg)

	config.Migrate(db)

	// SETUP
	uData := usrData.New(db)
	uSrv := usrSrv.New(uData, v, cld)
	uHdl := usrHdl.New(uSrv)

	hData := homeData.New(db)
	hSrv := homeSrv.New(hData, v, cld)
	hHdl := homeHdl.New(hSrv)

	fData := fbData.New(db)
	fSrv := fbSrv.New(fData)
	fHdl := fbHdl.New(fSrv)

	rData := rsvData.New(db)
	rSrv := rsvSrv.New(rData, v, coreapiMidtrans)
	rHdl := rsvHdl.New(rSrv)

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
	e.GET("/users", uHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", uHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", uHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/homestays", hHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/homestays", hHdl.List())
	e.GET("/homestays/:id", hHdl.GetbyID())
	e.PUT("/homestays/:id", hHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/homestays/:id", hHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/homestays", hHdl.Myhome(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/feedbacks", fHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/feedbacks", fHdl.List())
	e.GET("/feedbacks", fHdl.MyFeedback(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/reservations", rHdl.Create(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/reservations/callback", rHdl.Callback())
	e.GET("/reservations", rHdl.List(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/reservations/:id", rHdl.Detail(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/reservations/:id/accept", rHdl.Accept(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/reservations/:id/cancel", rHdl.Cancel(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
