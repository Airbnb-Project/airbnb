package main

import (
	"airbnb/config"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	// v := validator.New()
	// cld := helper.NewCloudinary(cfg)
	// coreapiMidtrans := helper.NewCoreMidtrans(cfg)

	config.Migrate(db)

	// MIDDLEWARE
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	// ROUTE

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
