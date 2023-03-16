package helper

import (
	"airbnb/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func ExtractToken(t interface{}) uint {
	user := t.(*jwt.Token)
	var userId uint
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userId = uint(claims["userID"].(float64))
		case int:
			userId = uint(claims["userID"].(int))
		}
	}
	return userId
}

func GenerateToken(id int) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = id
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 3 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, err := token.SignedString([]byte(config.JWT_KEY))
	if err != nil {
		log.Println(err.Error())
	}
	// log.Println(useToken, "/n", token)
	return useToken, token
}
