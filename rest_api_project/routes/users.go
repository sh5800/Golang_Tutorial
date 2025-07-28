package routes

import (
	"net/http"
	_ "strconv"

	"com.shreyash/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
		return
	}

	// event.UserID = 100

	err = user.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
		return
	}

	// data, err := json.Marshal(event)

	// if err != nil{
	// 	context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse your request","reason":err.Error()})
	// 	return
	// }

	context.JSON(http.StatusCreated,gin.H{"message":"User Created Successfully","user":user})
	
}