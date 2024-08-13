package api

import (
	"fmt"
	"net/http"

	"github.com/Efamamo/Event-Planning-System/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthService IUserUsecase
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	fmt.Println(u)

	c.IndentedJSON(http.StatusCreated, u)
	return
}
