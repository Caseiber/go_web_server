package stub

import (
	"fmt"
	"go_web_server/products"
)

// Stubbing out the store to allow for unit testing
type StubStore struct{}

func (s StubStore) GetProducts() ([]products.Product, error) {
	return []products.Product{}, nil
}

func (s StubStore) GetProduct(id string) (products.Product, error) {
	if id == fmt.Sprint(ErrorNoProduct) {
		return products.Product{}, products.ErrNoProduct
	}

	if id == fmt.Sprint(Success) {
		return ValidProduct, nil
	}

	return products.Product{}, nil
}

func (s StubStore) AddProduct(product products.Product) error {
	if product.ID == fmt.Sprint(ErrorBadProduct) {
		return products.ErrNoProduct
	}

	return nil
}

func (s StubStore) UpdateProduct(id string, product products.Product) ([]products.Product, error) {
	if id == fmt.Sprint(ErrorNoProduct) {
		return []products.Product{}, products.ErrNoProduct
	}

	return []products.Product{}, nil
}

func (s StubStore) DeleteProduct(id string) error { return nil }
