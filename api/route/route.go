package route

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	r.POST("/login")
	r.POST("/signup")
	r.POST("/logout")

	r.GET("/events")
	r.POST("/events")
	r.GET("/events/:id")
	r.PUT("/events/:id")
	r.DELETE("/events/:id")
}
