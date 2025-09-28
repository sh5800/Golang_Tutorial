package models

type Product struct{
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Description string `json:"description"`
	StockQuantity int64 `json:"stock_quantity"`
}