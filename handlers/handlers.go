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

type ProductService interface {
	ListProductsHandler(ps products.ProductStore) http.HandlerFunc
	CreateProductHandler(ps products.ProductStore) http.HandlerFunc
	GetProductHandler(ps products.ProductStore) http.HandlerFunc
	UpdateProductHandler(ps products.ProductStore) http.HandlerFunc
	DeleteProductHandler(ps products.ProductStore) http.HandlerFunc
}

type Service struct{}

func (s Service) ListProductsHandler(ps products.ProductStore) http.HandlerFunc {
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

func (s Service) CreateProductHandler(ps products.ProductStore) http.HandlerFunc {
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

func (s Service) GetProductHandler(ps products.ProductStore) http.HandlerFunc {
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

func (s Service) DeleteProductHandler(ps products.ProductStore) http.HandlerFunc {
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

func (s Service) UpdateProductHandler(ps products.ProductStore) http.HandlerFunc {
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
