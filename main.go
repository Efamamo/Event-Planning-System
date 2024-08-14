package main

import (
	"context"
	"log"

	controller "github.com/Efamamo/Event-Planning-System/api/controllers"
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
	eventRepo := repositories.NewEventsRepo(client)

	passwordService := infrastructure.PasswordService{}
	jwtService := infrastructure.Token{}

	authUsecase := usecases.AuthUsecase{AuthRepo: authRepo, PasswordService: passwordService, JWTService: jwtService}
	authController := controller.AuthController{AuthService: authUsecase}

	eventUsecase := usecases.EventsService{EventsRepo: eventRepo, AuthRepo: authRepo, JWTService: jwtService}
	eventController := controller.EventsController{EventsService: eventUsecase}

	route.StartServer(authController, eventController)

}
