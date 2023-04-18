package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go_web_server/products"

	"github.com/gorilla/mux"
)

func ListProducts(ps products.ProductStore) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productList, err := ps.GetProducts()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)

		data, err := json.Marshal(productList)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Write(data)
	}
}

func CreateProduct(ps products.ProductStore) http.HandlerFunc {
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

		err = ps.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
	}

}

func GetProductByID(ps products.ProductStore) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productID := mux.Vars(r)["id"]

		product, err := ps.GetProduct(productID)
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

func DeleteProductByID(ps products.ProductStore) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productID := mux.Vars(r)["id"]

		err := ps.DeleteProduct(productID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusAccepted)
		rw.Header().Add("content-type", "application/json")
	}
}

func UpdateProductByID(ps products.ProductStore) http.HandlerFunc {
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

		updatedProducts, err := ps.UpdateProduct(productID, updatedProduct)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusAccepted)
		rw.Header().Add("content-type", "application/json")
		marshaledData, err := json.Marshal(updatedProducts)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Write(marshaledData)
	}
}
