package main

import (
	"com.shreyash/rest-api/db"
	"com.shreyash/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server:= gin.Default() //provisioning a server
	routes.RegisterRoutes(server)
	server.Run(":8090")  //configuring a port on which this server will run
}

