package api

import (
	"fmt"
	"net/http"

	"github.com/Efamamo/Event-Planning-System/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService IUserUsecase
}

func (ac AuthController) Signup(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
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
