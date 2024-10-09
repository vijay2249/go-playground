package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wedonttrack.org/gorest/constants"
	"wedonttrack.org/gorest/utils"
)

func AuthenticateRequest(context *gin.Context) {
	authToken := context.Request.Header.Get("Authorization")
	if authToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := utils.VerifyToken(authToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	context.Set(constants.USER_ID, userId)
	context.Next()
}