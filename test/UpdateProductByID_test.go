package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_web_server/handlers"
	"go_web_server/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateProductByID(t *testing.T) {
	jsonProduct, err := json.Marshal(stub.ValidProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/products/"+fmt.Sprint(stub.Success), bytes.NewBuffer(jsonProduct))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.UpdateProductByID(store))
	handler.ServeHTTP(rr, req)
	checkResponseCode(http.StatusAccepted, rr.Code, t)
}
