package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wedonttrack.org/gorest/models"
	"wedonttrack.org/gorest/utils"
)

func user_signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func user_login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}
	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}