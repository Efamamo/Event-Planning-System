package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type IAuthRepository interface {
	Save(user domain.User) (*domain.User, error)
}
