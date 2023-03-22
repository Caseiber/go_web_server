package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"go_web_server/products"
)

var ErrNoProduct = errors.New("no product found")

func ListProducts() http.HandlerFunc {
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

func CreateProduct() http.HandlerFunc {
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

func GetProduct(id string) (products.Product, error) {
	data, err := products.GetProducts()
	if err != nil {
		return products.Product{}, err
	}

	var productList []products.Product
	err = json.Unmarshal(data, &productList)
	if err != nil {
		return products.Product{}, err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			return productList[i], nil
		}
	}

	return products.Product{}, ErrNoProduct
}

func UpdateProduct(id string, product products.Product) ([]byte, error) {
	data, err := products.GetProducts()
	if err != nil {
		return []byte{}, err
	}

	var productList []products.Product
	err = json.Unmarshal(data, &productList)
	if err != nil {
		return []byte{}, err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			productList[i] = product

			err = products.OverwriteProducts(productList)
			if err != nil {
				return []byte{}, err
			}

			data, err := products.GetProducts()
			if err != nil {
				return []byte{}, err
			}

			return data, nil
		}
	}

	return []byte{}, ErrNoProduct
}

func DeleteProduct(id string) error {
	data, err := products.GetProducts()
	if err != nil {
		return err
	}

	var productList []products.Product
	err = json.Unmarshal(data, &productList)
	if err != nil {
		return err
	}

	for i := 0; i < len(productList); i++ {

		if productList[i].ID == id {
			productList, err = removeElement(productList, i)
			if err != nil {
				return err
			}

			err = products.OverwriteProducts(productList)
			if err != nil {
				return err
			}
		}
	}

	return ErrNoProduct
}

func removeElement(productList []products.Product, index int) ([]products.Product, error) {
	if index < 0 || index >= len(productList) {
		return productList, errors.New("invalid index for deletion")
	}

	var updatedProducts []products.Product = make([]products.Product, 0)
	updatedProducts = append(updatedProducts, productList[0:index]...)
	updatedProducts = append(updatedProducts, productList[index+1:]...)

	return updatedProducts, nil
}
