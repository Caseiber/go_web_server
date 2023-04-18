package main

import (
	"fmt"
	"go_web_server/handlers"
	"go_web_server/products"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/products", handlers.ListProducts(products.Store{})).Methods(("GET"))
	router.Handle("/products", handlers.CreateProduct(products.Store{})).Methods(("POST"))
	router.Handle("/products/{id}", handlers.GetProductByID(products.Store{})).Methods(("GET"))
	router.Handle("/products/{id}", handlers.DeleteProductByID(products.Store{})).Methods(("DELETE"))
	router.Handle("/products/{id}", handlers.UpdateProductByID(products.Store{})).Methods(("PUT"))

	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	fmt.Println("Product Catalog server running on Port 9090")

	server.ListenAndServe()
}
