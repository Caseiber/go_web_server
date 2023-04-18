package test

import (
	"go_web_server/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ListProducts(store))
	handler.ServeHTTP(rr, req)

	checkResponseCode(http.StatusFound, rr.Code, t)
}
