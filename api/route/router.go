package route

import (
	"github.com/Efamamo/Event-Planning-System/api"
	"github.com/Efamamo/Event-Planning-System/infrastructure"
	"github.com/gin-gonic/gin"
)

func StartServer(authController api.AuthController, eventController api.EventsController) {
	r := gin.Default()
	r.POST("/login", authController.Login)
	r.POST("/signup", authController.Signup)
	// r.POST("/logout")

	r.GET("/events", infrastructure.AuthMiddleware(), eventController.GetEvents)
	r.POST("/events", infrastructure.AuthMiddleware(), eventController.AddEvent)
	r.GET("/events/:id", infrastructure.AuthMiddleware(), eventController.GetEventById)
	r.PUT("/events/:id", infrastructure.AuthMiddleware(), eventController.UpdateEvent)
	r.DELETE("/events/:id", infrastructure.AuthMiddleware(), eventController.DeleteEvent)

	r.Run()
}
