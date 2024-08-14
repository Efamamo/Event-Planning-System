package interfaces


import "github.com/Efamamo/Event-Planning-System/domain"

type IEventsRepo interface {
	GetEvents(string) (*[]domain.Event, error)
	AddEvent(string, domain.Event) (*domain.Event, error)
	GetEventById(string) (*domain.Event, error)
	UpdateEvent(string, domain.Event) error
	DeleteEvent(string) error
}
