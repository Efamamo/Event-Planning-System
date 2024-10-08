package controller

import (
	"net/http"

	"github.com/Efamamo/Event-Planning-System/domain"
	"github.com/Efamamo/Event-Planning-System/usecases/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthService interfaces.IUserUsecase
}

func (ac AuthController) Signup(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()

			switch field {
			case "Username":
				errorMessages["username"] = "Username is required."
			case "Password":
				errorMessages["password"] = "Password is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	u, err := ac.AuthService.Signup(user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, u)

}

func (ac AuthController) Login(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()

			switch field {
			case "Username":
				errorMessages["username"] = "Username is required."
			case "Password":
				errorMessages["password"] = "Password is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	tok, err := ac.AuthService.Login(user)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"accessToken": tok})
}
