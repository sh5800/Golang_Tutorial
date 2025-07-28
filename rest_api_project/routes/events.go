package routes 

import (
	"net/http"
	"strconv"

	"com.shreyash/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEventHandler(context *gin.Context){
	events, err := models.GetEvents() 
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not process your request","reason":err.Error()})
		return
	} 

	context.JSON(http.StatusOK,events)
}

func getEvent(context *gin.Context){
	id, err := strconv.ParseInt(context.Param("id"),10,64) 
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Cannot find the given resource","reason":err.Error()})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Cannot find the given resource","reason":err.Error()})
		return
	}

	context.JSON(http.StatusOK,event)
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
		return
	}

	// event.UserID = 100

	err = event.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
		return
	}

	// data, err := json.Marshal(event)

	// if err != nil{
	// 	context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
	// 	return
	// }

	context.JSON(http.StatusCreated,gin.H{"message":"Event Created Successfully","event":event})
	
}

func updateEvent(context *gin.Context){
	id, err := strconv.ParseInt(context.Param("id"),10,64) 
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Cannot find the given resource","reason":err.Error()})
		return
	}

	_ , err = models.GetEventById(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch the given resource"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not Update the event","reason":err.Error()})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Event Updated Successfully","event":updatedEvent})
}

func deleteEvent(context *gin.Context){
	id, err := strconv.ParseInt(context.Param("id"),10,64) 
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Cannot find the given resource","reason":err.Error()})
		return
	}

	deletedEvent , err := models.GetEventById(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch the given resource"})
		return
	}

	err = deletedEvent.Delete()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not delete the given resource"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Event Deleted Successfully","event":deletedEvent})
}