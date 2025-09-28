package main

import (
	"net/http"
	"log"
	"com.shreyash/go_inventory/v2/pkg/routes"
	_"com.shreyash/go_inventory/v2/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

)

// @title Product Inventory API
// @version 1.0
// @description This is a sample server for managing products.
// @host localhost:8000
// @BasePath /
func main(){
	r :=  mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	routes.RegisterProductRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8000",r))
}