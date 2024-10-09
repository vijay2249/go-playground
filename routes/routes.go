package routes

import (
	"github.com/gin-gonic/gin"
	"wedonttrack.org/gorest/middlewares"
)

func RegisteredRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //for <param1> path - execute <param2> function
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.AuthenticateRequest, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", user_signup)
	server.POST("/login", user_login)
}