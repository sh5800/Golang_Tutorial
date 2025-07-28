package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getEventHandler) //adding a GET route with /events endpoint and adding a handler that would be executed whenever a request to /events is made 

	server.GET("/events/:id",getEvent)

	server.POST("/events",createEvent)

	server.PUT("/events/:id",updateEvent)

	server.DELETE("/events/:id",deleteEvent)

	server.POST("/signup",signup)
}