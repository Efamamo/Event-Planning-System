package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type AuthUsecase struct {
	AuthRepo IAuthRepository
}

func (au AuthUsecase) Signup(user domain.User) (*domain.User, error) {
	u, err := au.AuthRepo.Save(user)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (au AuthUsecase) Login(user domain.User) (string, error) {
	return "", nil
}
