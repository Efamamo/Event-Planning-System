package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

type Event struct {
	Name        string    `json:"name" bson:"name" binding:"required"`
	Description string    `json:"description" bson:"description" binding:"required"`
	Date        time.Time `json:"date" bson:"date" binding:"required"`
	Location    string    `json:"location" bson:"location" binding:"required"`
	Owner       string    `json:"owner" bson:"owner" binding:"required"`
}
