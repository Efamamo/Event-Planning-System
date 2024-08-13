package repositories

import (
	"context"
	"time"

	"github.com/Efamamo/Event-Planning-System/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	collection mongo.Collection
}

func NewUserRepo(client *mongo.Client) AuthRepository {
	userCollection := client.Database("event-management").Collection("users")

	return AuthRepository{
		collection: *userCollection,
	}
}

func (ar AuthRepository) Save(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()

	_, e := ar.collection.InsertOne(ctx, user)

	if e != nil {
		return nil, e
	}

	return &user, nil
}
