package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"wedonttrack.org/gorest/constants"
	"wedonttrack.org/gorest/models"
)

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect input data"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldnt fetch data"})
		return
	}
	context.JSON(http.StatusOK, event)
}


func getEvents(context *gin.Context) {
	fmt.Println(context)
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events data"})
		return
	}
	context.JSON(http.StatusOK, events)

	// send json response with response status code and json content
	// context.JSON(http.StatusOK, gin.H{"message": "Hi Mom!"})
}

func createEvent (context *gin.Context){

	
	var event models.Event
	err := context.ShouldBindJSON(&event) //datatype of the input(request body) should be same as the event variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}
	event.UserID = context.GetInt64(constants.USER_ID)
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect input data"})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(400, gin.H{"message": "reading event data went wrong"})
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse input data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect input data"})
		return
	}

	err = models.DeleteEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "unable to delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "deleted event"})
}