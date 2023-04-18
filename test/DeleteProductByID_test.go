package test

import (
	"fmt"
	"go_web_server/handlers"
	"go_web_server/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteProductByID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/products/"+fmt.Sprint(stub.Success), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.DeleteProductByID(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusAccepted, rr.Code, t)
}
