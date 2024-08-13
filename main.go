package main

import (
	"github.com/Efamamo/Event-Planning-System/api"
	"github.com/Efamamo/Event-Planning-System/api/route"
	"github.com/Efamamo/Event-Planning-System/usecases"
)

func main() {

	authUsecase := usecases.AuthUsecase{}
	authController := api.AuthController{AuthService: authUsecase}
	route.StartServer(authController)

}
