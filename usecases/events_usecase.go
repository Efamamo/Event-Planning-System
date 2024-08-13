package usecases

import "github.com/Efamamo/Event-Planning-System/domain"

type EventsService struct {
	EventsRepo IEventsRepo
}

func (es EventsService) GetEvents(username string) (*[]domain.Event, error) {
	ev, err := es.EventsRepo.GetEvents(username)

	if err != nil {
		return nil, err
	}

	return ev, nil
}
