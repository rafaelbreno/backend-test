package controllers

import (
	"coderockr-test/app/helpers"
	"coderockr-test/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (_ UserController) SignUp(c *gin.Context) {
	var userInput models.CreateUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userInput.Password != userInput.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both password are different"})
		return
	}

	password, err := helpers.HashPassword(userInput.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when hashing password, contact the administrator"})
		return
	}

	user := models.User{
		Name:           userInput.Name,
		Email:          userInput.Email,
		Password:       string(password),
		Biography:      userInput.Biography,
		ProfilePicture: userInput.ProfilePicture,
		City:           userInput.City,
		State:          userInput.State,
	}

	models.DB.Model(&user)

	c.JSON(http.StatusOK, gin.H{"data": &user})
}
