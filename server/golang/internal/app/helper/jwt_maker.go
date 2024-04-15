package helper

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Token struct {
	secret_key string
}

type TokenPayload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int32     `json:"user_id"`
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

func (t *Token) CreateToken(email string, user_id int32, role string, duration time.Duration) (string, *TokenPayload, error) {
	payload := &TokenPayload{
		ID:        uuid.New(),
		UserID:    user_id,
		Role:      role,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(t.secret_key))

	if err != nil {
		fmt.Println("error signing token: ", err)
	}
	return tokenString, payload, err
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
		return nil, err
	}
	claims, ok := jwt_token.Claims.(*TokenPayload)

	expired := claims.Valid()
	if !ok {
		return nil, ErrInvalidToken
	}
	if expired != nil {
		return nil, expired
	}

	if ok {
		fmt.Println(claims.Email, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return claims, err
}
