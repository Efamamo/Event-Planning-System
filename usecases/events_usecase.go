package usecases

import (
	"errors"

	"github.com/Efamamo/Event-Planning-System/domain"
	"github.com/Efamamo/Event-Planning-System/usecases/interfaces"
)

type EventsService struct {
	EventsRepo interfaces.IEventsRepo
	AuthRepo   interfaces.IAuthRepository
	JWTService interfaces.IJWTService
}

func (es EventsService) GetEvents(authPart string) (*[]domain.Event, error) {

	token, err := es.JWTService.ValidateToken(authPart)
	if err != nil {
		return nil, errors.New("unauthorized")
	}
	username, err := es.JWTService.GetUserName(token)

	if err != nil {
		return nil, errors.New("unauthorized")
	}

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

func (es EventsService) GetEventById(id string, authPart string) (*domain.Event, error) {
	token, err := es.JWTService.ValidateToken(authPart)
	if err != nil {
		return nil, err
	}
	username, err := es.JWTService.GetUserName(token)
	if err != nil {
		return nil, err
	}

	ev, err := es.EventsRepo.GetEventById(id)
	if err != nil {
		return nil, err
	}

	if ev.Owner != username {
		return nil, errors.New("unauthorized")
	}

	return ev, nil
}
func (es EventsService) AddEvent(authPart string, event domain.Event) (*domain.Event, error) {
	token, err := es.JWTService.ValidateToken(authPart)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	username, err := es.JWTService.GetUserName(token)

	if err != nil {
		return nil, errors.New("unauthorized")
	}

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

func (es EventsService) UpdateEvent(id string, event domain.Event) error {

	err := es.EventsRepo.UpdateEvent(id, event)

	if err != nil {
		return err
	}
	return nil
}

func (es EventsService) DeleteEvent(id string, authPart string) error {
	token, err := es.JWTService.ValidateToken(authPart)
	if err != nil {
		return errors.New("unauthorized")
	}

	username, err := es.JWTService.GetUserName(token)
	if err != nil {
		return err
	}

	ev, err := es.EventsRepo.GetEventById(id)

	if err != nil {
		return errors.New("unauthorized")
	}

	if ev.Owner != username {
		return errors.New("unauthorized")
	}

	err = es.EventsRepo.DeleteEvent(id)
	if err != nil {
		return err
	}

	return nil

}

func (es EventsService) CheckValidity(id string, authPart string) error {
	token, err := es.JWTService.ValidateToken(authPart)
	if err != nil {
		return err
	}

	username, err := es.JWTService.GetUserName(token)
	if err != nil {
		return err
	}

	ev, err := es.EventsRepo.GetEventById(id)
	if err != nil {
		return err
	}

	if ev.Owner != username {
		return errors.New("unauthorized")
	}
	return nil
}
