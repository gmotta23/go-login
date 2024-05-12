package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(ID uint) (string, error) {
	mySigningKey := []byte("JWTKEY")

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3600)),
		ID:        strconv.FormatUint(uint64(ID), 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}
