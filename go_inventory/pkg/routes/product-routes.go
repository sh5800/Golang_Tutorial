package routes

import (
	"github.com/gorilla/mux"
	"com.shreyash/go_inventory/v2/pkg/controllers"
)

var RegisterProductRoutes = func(r *mux.Router){
	r.HandleFunc("/products",controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}",controllers.GetProduct).Methods("GET")
	r.HandleFunc("/products",controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}/increase",controllers.IncreaseStockQuantity).Methods("PUT")
	r.HandleFunc("/products/{id}/decrease",controllers.DecreaseStockQuantity).Methods("PUT")
	r.HandleFunc("/products/{id}",controllers.DeleteProduct).Methods("DELETE")
}