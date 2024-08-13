package infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordService struct{}

func (ps PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil
}
