package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}
