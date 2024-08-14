package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
}

func (tok Token) ValidateToken(t string) (*jwt.Token, error) {
	token, e := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("JwtSecret"), nil
	})

	if e != nil || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	return token, nil
}

func (tok Token) GetUserName(token *jwt.Token) (string, error) {

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", errors.New("Unauthorized")
	}

	username, ok := claims["username"].(string)
	if !ok || username == "" {
		return "", errors.New("Unauthorized")
	}
	return username, nil
}

func (tok Token) GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString([]byte("JwtSecret"))

	if e != nil {
		return "", errors.New("Cant Sign Token")
	}

	return jwtToken, nil
}
