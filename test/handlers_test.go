package test

import (
	"go_web_server/handlers"
	"go_web_server/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

var store = stub.StubStore{}

func TestListProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ListProducts(store))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v, want %v",
			status, http.StatusFound)
	}
}

func TestGetProductByID(t *testing.T) {
	id := "0"
	req, err := http.NewRequest("GET", "/products/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetProductByID(store))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}
}
