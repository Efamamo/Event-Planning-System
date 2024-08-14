package usecases

import (
	"errors"

	"github.com/Efamamo/Event-Planning-System/domain"
)

type EventsService struct {
	EventsRepo IEventsRepo
	AuthRepo   IAuthRepository
}

func (es EventsService) GetEvents(username string) (*[]domain.Event, error) {

	user, err := es.AuthRepo.FindUser(username)

	if err != nil {
		return nil, err
	}

	ev, err := es.EventsRepo.GetEvents(user.Username)

	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (es EventsService) GetEventById(id string, username string) (*domain.Event, error) {
	ev, err := es.EventsRepo.GetEventById(id)

	if err != nil {
		return nil, err
	}

	if ev.Owner != username {
		return nil, errors.New("unauthorized")
	}

	return ev, nil
}
func (es EventsService) AddEvent(username string, event domain.Event) (*domain.Event, error) {

	user, err := es.AuthRepo.FindUser(username)

	if err != nil {
		return nil, err
	}

	event.Owner = user.Username
	ev, err := es.EventsRepo.AddEvent(user.Username, event)

	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (es EventsService) UpdateEvent(id string, event domain.Event, username string) error {

	e, err := es.EventsRepo.GetEventById(id)
	if err != nil {
		return err
	}

	if e.Owner != username {
		return errors.New("unauthorized")
	}

	err = es.EventsRepo.UpdateEvent(id, event)

	if err != nil {
		return err
	}
	return nil
}

func (es EventsService) DeleteEvent(id string, username string) error {
	e, err := es.EventsRepo.GetEventById(id)

	if err != nil {
		return err
	}
	if e.Owner != username {
		return errors.New("unauthorized")
	}

	err = es.EventsRepo.DeleteEvent(id)

	if err != nil {
		return err
	}
	return nil

}
