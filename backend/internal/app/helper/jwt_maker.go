package helper

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Token struct {
	secret_key string
}

type TokenPayload struct {
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func (payload *TokenPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewTokenMaker(secretkey string) *Token {
	return &Token{
		secret_key: secretkey,
	}
}

func (t *Token) CreateToken(email string, role string) (string, error) {
	new_payload := &TokenPayload{
		Role:      role,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(60 * time.Minute),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, new_payload)
	tokenString, err := token.SignedString([]byte(t.secret_key))

	if err != nil {
		fmt.Println("error signing token: ", err)
	}
	return tokenString, err
}

func (t *Token) VerifyToken(token string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(t.secret_key), nil
	}
	jwt_token, err := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)
	if err != nil {
		fmt.Println(err)
	}
	claims, ok := jwt_token.Claims.(*TokenPayload)
	if ok {
		fmt.Println(claims.Email, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return claims, err
}
