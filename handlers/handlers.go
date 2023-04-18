package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"go_web_server/products"

	"github.com/gorilla/mux"
)

func ListProducts(ps products.ProductStore) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		productList, err := ps.GetProducts()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(productList)
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

		ps.AddProduct(product)
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
		productID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			return
		}

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
		productID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			return
		}

		err = ps.DeleteProduct(productID)
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

		productID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			return
		}

		updatedProducts, err := ps.UpdateProduct(productID, updatedProduct)
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
