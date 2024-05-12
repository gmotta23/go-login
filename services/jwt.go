package services

import (
	"gmtc/login/utils"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(ID uint) (string, error) {
	mySigningKey := []byte("JWTKEY")
	tokenExpirationTime := time.Hour * 24

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpirationTime)),
		ID:        utils.UintToString(ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateAndParseJWTToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWTKEY"), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		return claims, nil
	} else {
		log.Fatal("JWT parsing failed")
		return nil, nil
	}
}
