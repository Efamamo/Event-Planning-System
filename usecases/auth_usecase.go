package usecases

import (
	"github.com/Efamamo/Event-Planning-System/domain"
)

type AuthUsecase struct {
	AuthRepo        IAuthRepository
	PasswordService IPassword
	JWTService      IJWTService
}

func (au AuthUsecase) Signup(user domain.User) (*domain.User, error) {
	hashedPassword, e := au.PasswordService.HashPassword(user.Password)

	if e != nil {
		return nil, e
	}
	user.Password = hashedPassword
	user.Events = []domain.Event{}
	u, err := au.AuthRepo.Save(user)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (au AuthUsecase) Login(user domain.User) (string, error) {
	u, err := au.AuthRepo.FindUser(user.Username)

	if err != nil {
		return "", err
	}

	_, err = au.PasswordService.ComparePassword(u.Password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := au.JWTService.GenerateToken(u.Username)

	return token, nil

}
