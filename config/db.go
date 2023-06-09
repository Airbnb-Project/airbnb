package config

import (
	feedback "airbnb/feature/feedback/data"
	home "airbnb/feature/homestay/data"
	reservation "airbnb/feature/reservation/data"
	"airbnb/feature/user/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&home.Homestay{})
	db.AutoMigrate(&home.Image{})
	db.AutoMigrate(&feedback.Feedback{})
	db.AutoMigrate(&reservation.Reservation{})

}
