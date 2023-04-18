package test

import (
	"fmt"
	"go_web_server/handlers"
	"go_web_server/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/"+fmt.Sprint(stub.Success), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetProductByID(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusFound, rr.Code, t)
}
