package repositories

import (
	"context"
	"time"

	"github.com/Efamamo/Event-Planning-System/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventsRepository struct {
	collection mongo.Collection
}

func NewEventsRepo(client *mongo.Client) EventsRepository {
	userCollection := client.Database("event-management").Collection("events")

	return EventsRepository{
		collection: *userCollection,
	}
}

func (er EventsRepository) GetEventById(id string) (*domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": Id}
	var event domain.Event
	err = er.collection.FindOne(ctx, filter).Decode(&event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (er EventsRepository) GetEvents(username string) (*[]domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"owner": username}
	cur, err := er.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	events := make([]domain.Event, 0)

	for cur.Next(ctx) {
		var event domain.Event
		err := cur.Decode(&event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}

	return &events, nil

}

func (er EventsRepository) AddEvent(username string, event domain.Event) (*domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.ID = primitive.NewObjectID()

	_, e := er.collection.InsertOne(ctx, event)

	if e != nil {
		return nil, e
	}

	return &event, nil
}

func (er EventsRepository) UpdateEvent(id string, uevent domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filterId := bson.M{"_id": Id}
	filter := bson.M{
		"$set": bson.M{
			"name":        uevent.Name,
			"description": uevent.Description,
			"date":        uevent.Date,
			"location":    uevent.Location,
		},
	}

	_, err = er.collection.UpdateOne(ctx, filterId, filter)

	if err != nil {
		return err
	}
	return nil
}

func (er EventsRepository) DeleteEvent(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": Id}

	_, err = er.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
