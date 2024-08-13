package repositories

import (
	"context"
	"time"

	"github.com/Efamamo/Event-Planning-System/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventsRepository struct {
	collection mongo.Collection
}

func NewEventsRepo(client *mongo.Client) EventsRepository {
	userCollection := client.Database("event-management").Collection("users")

	return EventsRepository{
		collection: *userCollection,
	}
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
