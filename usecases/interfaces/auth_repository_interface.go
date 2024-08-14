package interfaces

import "github.com/Efamamo/Event-Planning-System/domain"

type IAuthRepository interface {
	Save(user domain.User) (*domain.User, error)
	FindUser(username string) (*domain.User, error)
}
