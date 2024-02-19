package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24)

	tokenString, err := token.SignedString([]byte("hung1324"))

	if err != nil {
		fmt.Println("error signing token: ", err)
	}
	return tokenString, err
}
