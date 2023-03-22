package main

import (
	"fmt"
	"go_web_server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/products", handlers.ListProducts()).Methods(("GET"))
	router.Handle("/products", handlers.CreateProduct()).Methods(("POST"))

	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	fmt.Println("Product Catalog server running on Port 9090")

	server.ListenAndServe()
}
