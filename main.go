package main

import (
	"github.com/gin-gonic/gin"
	"wedonttrack.org/gorest/db"
	"wedonttrack.org/gorest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisteredRoutes(server)
	server.Run(":8080")
}

