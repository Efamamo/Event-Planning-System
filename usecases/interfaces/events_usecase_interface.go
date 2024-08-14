package interfaces

import "github.com/Efamamo/Event-Planning-System/domain"

type IEventsService interface {
	CheckValidity(id string, username string) error
	GetEvents(string) (*[]domain.Event, error)
	AddEvent(string, domain.Event) (*domain.Event, error)
	UpdateEvent(string, domain.Event) error
	GetEventById(id string, username string) (*domain.Event, error)
	DeleteEvent(id string, username string) error
}
