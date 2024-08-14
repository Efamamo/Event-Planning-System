package route

import (
	controller "github.com/Efamamo/Event-Planning-System/api/controllers"
	"github.com/Efamamo/Event-Planning-System/infrastructure"
	"github.com/gin-gonic/gin"
)

func StartServer(authController controller.AuthController, eventController controller.EventsController) {
	r := gin.Default()
	r.POST("/login", authController.Login)
	r.POST("/signup", authController.Signup)

	r.GET("/events", infrastructure.AuthMiddleware(), eventController.GetEvents)
	r.POST("/events", infrastructure.AuthMiddleware(), eventController.AddEvent)
	r.GET("/events/:id", infrastructure.AuthMiddleware(), eventController.GetEventById)
	r.PUT("/events/:id", infrastructure.AuthMiddleware(), eventController.UpdateEvent)
	r.DELETE("/events/:id", infrastructure.AuthMiddleware(), eventController.DeleteEvent)

	r.Run()
}
