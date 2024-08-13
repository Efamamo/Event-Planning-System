package route

import (
	"github.com/Efamamo/Event-Planning-System/api"
	"github.com/gin-gonic/gin"
)

func StartServer(authController api.AuthController) {
	r := gin.Default()
	r.POST("/login")
	r.POST("/signup", authController.Signup)
	// r.POST("/logout")

	r.GET("/events")
	r.POST("/events")
	r.GET("/events/:id")
	r.PUT("/events/:id")
	r.DELETE("/events/:id")

	r.Run()
}
