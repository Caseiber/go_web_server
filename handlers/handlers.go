package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go_web_server/products"

	"github.com/gorilla/mux"
)

func ListProductsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productList, err := products.GetProducts()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(productList)
	}
}

func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var product products.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Could not unmarshal data"))
			return
		}

		products.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
	}

}

func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productID := mux.Vars(r)["id"]
		product, err := products.GetProduct(productID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusFound)
		rw.Header().Add("content-type", "application/json")
		rw.Write(data)
	}
}

func DeleteProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productID := mux.Vars(r)["id"]
		err := products.DeleteProduct(productID)
		fmt.Println(err)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusAccepted)
		rw.Header().Add("content-type", "application/json")
	}
}

func UpdateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var updatedProduct products.Product
		err = json.Unmarshal(data, &updatedProduct)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			return
		}

		productID := mux.Vars(r)["id"]
		updatedProducts, err := products.UpdateProduct(productID, updatedProduct)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusAccepted)
		rw.Header().Add("content-type", "application/json")
		rw.Write(updatedProducts)
	}
}
