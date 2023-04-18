package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_web_server/handlers"
	"go_web_server/products"
	"go_web_server/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	jsonProduct, err := json.Marshal(stub.ValidProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonProduct))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateProduct(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusCreated, rr.Code, t)
}

func TestCreateProductCannotUnmarshal(t *testing.T) {
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte("not a product")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateProduct(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusExpectationFailed, rr.Code, t)
}

func TestCreateProductCannotAdd(t *testing.T) {
	var invalidProduct = products.Product{
		ID: fmt.Sprint(stub.ErrorBadProduct),
	}

	jsonProduct, err := json.Marshal(invalidProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonProduct))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateProduct(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusInternalServerError, rr.Code, t)
}
