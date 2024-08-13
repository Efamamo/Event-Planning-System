package route

import (
	"github.com/Efamamo/Event-Planning-System/api"
	"github.com/gin-gonic/gin"
)

func StartServer(authController api.AuthController, eventController api.EventsController) {
	r := gin.Default()
	r.POST("/login", authController.Login)
	r.POST("/signup", authController.Signup)
	// r.POST("/logout")

	r.GET("/events", eventController.GetEvents)
	r.POST("/events")
	r.GET("/events/:id")
	r.PUT("/events/:id")
	r.DELETE("/events/:id")

	r.Run()
}
