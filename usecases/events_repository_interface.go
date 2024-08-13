package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type IEventsRepo interface {
	GetEvents(string) (*[]domain.Event, error)
}
