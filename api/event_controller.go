package api

import (
	"net/http"
	"strings"

	"github.com/Efamamo/Event-Planning-System/usecases"
	"github.com/gin-gonic/gin"
)

type EventsController struct {
	EventsService IEventsService
	JWTService    usecases.IJWTService
}

func (ec EventsController) GetEvents(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Header("Content-Type", "application/json")
		c.IndentedJSON(401, gin.H{"message": "unauthorized"})

		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.Header("Content-Type", "application/json")
		c.IndentedJSON(401, gin.H{"message": "unauthorized"})

		return
	}

	username, err := ec.JWTService.GetUserName(authParts[1])

	if err != nil {
		c.IndentedJSON(401, gin.H{"message": "unauthorized"})
		return
	}

	events, err := ec.EventsService.GetEvents(username)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}
