package infrastructure

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func (ps PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil
}

func (ps PasswordService) ComparePassword(euPassword string, uPassword string) (bool, error) {

	if bcrypt.CompareHashAndPassword([]byte(euPassword), []byte(uPassword)) != nil {
		return false, errors.New("Invalid Credentials")
	}

	return true, nil
}

