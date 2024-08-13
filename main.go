package main

import (
	"context"
	"log"

	"github.com/Efamamo/Event-Planning-System/api"
	"github.com/Efamamo/Event-Planning-System/api/route"
	"github.com/Efamamo/Event-Planning-System/infrastructure"
	"github.com/Efamamo/Event-Planning-System/infrastructure/repositories"
	"github.com/Efamamo/Event-Planning-System/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	authRepo := repositories.NewUserRepo(client)
	passwordService := infrastructure.PasswordService{}
	jwtService := infrastructure.Token{}
	authUsecase := usecases.AuthUsecase{AuthRepo: authRepo, PasswordService: passwordService, JWTService: jwtService}
	authController := api.AuthController{AuthService: authUsecase}


	eventRepo := repositories.NewEventsRepo(client)
	eventUsecase := usecases.EventsService{EventsRepo: eventRepo}
	eventController := api.EventsController{EventsService: eventUsecase,JWTService: jwtService}

	route.StartServer(authController, eventController)

}
