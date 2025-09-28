package controllers

import(
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"math/rand"
	"com.shreyash/go_inventory/v2/pkg/models"
)

var products []models.Product

// GetProducts godoc
// @Summary List all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProduct godoc
// @Summary Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {string} string "Product not found"
// @Router /products/{id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range products{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w,"Product Not Found",http.StatusNotFound)
}

// DeleteProduct godoc
// @Summary Delete product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {string} string "Product not found"
// @Router /products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range products{
		if item.ID == params["id"]{
			products = append(products[:index],products[index+1:]...)
			json.NewEncoder(w).Encode(products)
			return
		}
	}
	http.Error(w,"Product Not Found",http.StatusNotFound)
}

// CreateProduct godoc
// @Summary Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product to create"
// @Success 200 {object} models.Product
// @Router /products [post]
func CreateProduct(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa((rand.Intn(1001)))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// IncreaseStockQuantity godoc
// @Summary Increase Stock Quantity for a Product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id}/increase [put]
func IncreaseStockQuantity(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range products{
		if item.ID == params["id"]{

			var product models.Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			products[index].StockQuantity += product.StockQuantity
			json.NewEncoder(w).Encode(products[index])
			return
		}
	}
	http.Error(w,"Product Not Found",http.StatusNotFound)
}

// DecreaseStockQuantity godoc
// @Summary Decrease Stock Quantity for a Product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id}/decrease [put]
func DecreaseStockQuantity(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range products{
		if item.ID == params["id"]{

			var product models.Product
			_ = json.NewDecoder(r.Body).Decode(&product)

			if(products[index].StockQuantity - product.StockQuantity < 0){
				http.Error(w,"Stock Quantity Will be Negative",http.StatusBadRequest)
				return
			}
			products[index].StockQuantity -= product.StockQuantity
			json.NewEncoder(w).Encode(products[index])
			return
		}
	}
	http.Error(w,"Product Not Found",http.StatusNotFound)
}