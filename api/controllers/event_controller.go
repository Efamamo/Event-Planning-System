package controller

import (
	"net/http"
	"strings"

	"github.com/Efamamo/Event-Planning-System/domain"
	"github.com/Efamamo/Event-Planning-System/usecases/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventsController struct {
	EventsService interfaces.IEventsService
}

func (ec EventsController) GetEvents(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	events, err := ec.EventsService.GetEvents(authParts[1])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func (ec EventsController) GetEventById(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	ev, err := ec.EventsService.GetEventById(id, authParts[1])
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ev)
}
func (ec EventsController) AddEvent(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	var event domain.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()

			switch field {
			case "Name":
				errorMessages["name"] = "name is required."
			case "Description":
				errorMessages["description"] = "description is required."
			case "Location":
				errorMessages["location"] = "location is required."
			case "Date":
				errorMessages["data"] = "date is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	ev, err := ec.EventsService.AddEvent(authParts[1], event)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, ev)

}

func (ec EventsController) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	e := ec.EventsService.CheckValidity(id, authParts[1])
	if e != nil {
		if e.Error() == "unauthorized" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Id Format"})
			return
		}
	}

	var event domain.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()

			switch field {
			case "Name":
				errorMessages["name"] = "name is required."
			case "Description":
				errorMessages["description"] = "description is required."
			case "Location":
				errorMessages["location"] = "location is required."
			case "Date":
				errorMessages["date"] = "date is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	e = ec.EventsService.UpdateEvent(id, event)

	if e != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	}

}

func (ec EventsController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	err := ec.EventsService.DeleteEvent(id, authParts[1])

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}
