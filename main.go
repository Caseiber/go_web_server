package main

import (
	"fmt"
	"go_web_server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/products", handlers.ListProductsHandler()).Methods(("GET"))
	router.Handle("/products", handlers.CreateProductHandler()).Methods(("POST"))
	router.Handle("/products/{id}", handlers.GetProductHandler()).Methods(("GET"))
	router.Handle("/products/{id}", handlers.DeleteProductHandler()).Methods(("DELETE"))
	router.Handle("/products/{id}", handlers.UpdateProductHandler()).Methods(("PUT"))

	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	fmt.Println("Product Catalog server running on Port 9090")

	server.ListenAndServe()
}
