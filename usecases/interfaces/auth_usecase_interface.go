package interfaces

import "github.com/Efamamo/Event-Planning-System/domain"

type IUserUsecase interface {
	Signup(user domain.User) (*domain.User, error)
	Login(user domain.User) (string, error)
}
