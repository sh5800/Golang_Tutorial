package controllers

import (
	_"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gorilla/mux"
	"com.shreyash/go_inventory/v2/pkg/models"
)

// resetProducts is a helper to reset the products slice before each test
func resetProducts() {
	products = []models.Product{
		{ID: "1", Name: "Soap", Description: "Dettol Soap", StockQuantity: 3},
		{ID: "2", Name: "Oil", Description: "Fortune Olive Oil", StockQuantity: 10},
	}
}

// ------------------------
// GetProducts Tests
// ------------------------
func TestGetProducts(t *testing.T) {
	resetProducts()

	req := httptest.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}

	var got []models.Product
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("expected 2 products, got %d", len(got))
	}
}

// ------------------------
// GetProduct Tests
// ------------------------
func TestGetProduct(t *testing.T) {
	resetProducts()

	tests := []struct {
		id           string
		expectedCode int
	}{
		{"1", http.StatusOK},
		{"2", http.StatusOK},
		{"100", http.StatusNotFound}, // non-existent product
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/products/"+tt.id, nil)
		rr := httptest.NewRecorder()
		req = muxSetVars(req, map[string]string{"id": tt.id})
		handler := http.HandlerFunc(GetProduct)
		handler.ServeHTTP(rr, req)

		if rr.Code != tt.expectedCode {
			t.Errorf("GetProduct(%s): expected %v, got %v", tt.id, tt.expectedCode, rr.Code)
		}
	}
}

// ------------------------
// CreateProduct Tests
// ------------------------
func TestCreateProduct(t *testing.T) {
	resetProducts()

	body := `{"name":"Notebook","description":"Classmate Spiral Notebook","stock_quantity":50}`
	req := httptest.NewRequest("POST", "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}

	var p models.Product
	if err := json.Unmarshal(rr.Body.Bytes(), &p); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if p.Name != "Notebook" || p.StockQuantity != 50 {
		t.Errorf("unexpected product returned: %+v", p)
	}
}

// ------------------------
// DeleteProduct Tests
// ------------------------
func TestDeleteProduct(t *testing.T) {
	resetProducts()

	tests := []struct {
		id           string
		expectedCode int
		expectedLen  int
	}{
		{"1", http.StatusOK, 1},
		{"100", http.StatusNotFound, 2}, // deleting non-existent
	}

	for _, tt := range tests {
		req := httptest.NewRequest("DELETE", "/products/"+tt.id, nil)
		rr := httptest.NewRecorder()
		req = muxSetVars(req, map[string]string{"id": tt.id})
		handler := http.HandlerFunc(DeleteProduct)
		handler.ServeHTTP(rr, req)

		if rr.Code != tt.expectedCode {
			t.Errorf("DeleteProduct(%s): expected %v, got %v", tt.id, tt.expectedCode, rr.Code)
		}

		if len(products) != tt.expectedLen {
			t.Errorf("DeleteProduct(%s): expected products length %d, got %d", tt.id, tt.expectedLen, len(products))
		}
	}
}

// ------------------------
// IncreaseStock Tests
// ------------------------
func TestIncreaseStockQuantity(t *testing.T) {
	resetProducts()

	body := `{"stock_quantity":5}`
	req := httptest.NewRequest("PUT", "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	req = muxSetVars(req, map[string]string{"id": "1"})
	handler := http.HandlerFunc(IncreaseStockQuantity)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}

	if products[0].StockQuantity != 8 { // 3+5
		t.Errorf("expected stock 8, got %d", products[0].StockQuantity)
	}
}

// ------------------------
// DecreaseStock Tests
// ------------------------
func TestDecreaseStockQuantity(t *testing.T) {
	resetProducts()

	tests := []struct {
		id           string
		body         string
		expectedCode int
		expectedQty  int64
	}{
		{"1", `{"stock_quantity":2}`, http.StatusOK, 1},  // 3-2=1
		{"1", `{"stock_quantity":5}`, http.StatusBadRequest, 3}, // cannot go negative
	}

	for _, tt := range tests {
		req := httptest.NewRequest("PUT", "/products/"+tt.id, strings.NewReader(tt.body))
		req.Header.Set("Content-Type", "application/json")
		req = muxSetVars(req, map[string]string{"id": tt.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(DecreaseStockQuantity)
		handler.ServeHTTP(rr, req)

		if rr.Code != tt.expectedCode {
			t.Errorf("DecreaseStockQuantity(%s): expected %v, got %v", tt.id, tt.expectedCode, rr.Code)
		}

		if tt.expectedCode == http.StatusOK && products[0].StockQuantity != tt.expectedQty {
			t.Errorf("expected stock %d, got %d", tt.expectedQty, products[0].StockQuantity)
		}
	}
}

// ------------------------
// Helper to set mux vars in tests
// ------------------------
func muxSetVars(r *http.Request, vars map[string]string) *http.Request {
	return mux.SetURLVars(r, vars)
}
