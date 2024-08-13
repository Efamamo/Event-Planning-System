package api

import "github.com/Efamamo/Event-Planning-System/domain"

type IEventsService interface {
	GetEvents(string) (*[]domain.Event, error)
}
