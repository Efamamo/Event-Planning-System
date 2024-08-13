package infrastructure

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct{}

func (tok Token) ValidateToken(t string) (*jwt.Token, error) {
	token, e := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JwtSecret")), nil
	})

	if e != nil || !token.Valid {
		return nil, errors.New("Unauthorized")
	}
	return token, nil
}

func (tok Token) GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString([]byte(os.Getenv("JwtSecret")))

	if e != nil {
		return "", errors.New("Cant Sign Token")
	}

	return jwtToken, nil
}
