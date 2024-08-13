package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type AuthUsecase struct {
	AuthRepo        IAuthRepository
	PasswordService IPassword
}

func (au AuthUsecase) Signup(user domain.User) (*domain.User, error) {
	hashedPassword, e := au.PasswordService.HashPassword(user.Password)

	if e != nil {
		return nil, e
	}
	user.Password = hashedPassword
	u, err := au.AuthRepo.Save(user)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (au AuthUsecase) Login(user domain.User) (string, error) {
	return "", nil
}
