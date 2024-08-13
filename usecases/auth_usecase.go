package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type AuthUsecase struct{}

func (au AuthUsecase) Signup(user domain.User) (*domain.User, error) {
	return &user, nil
}

func (au AuthUsecase) Login(user domain.User) (string, error) {
	return "", nil
}
