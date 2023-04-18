package products

import (
	"encoding/json"
	"errors"
	"os"
)

const DataLocation = "./data/data.json"

var ErrNoProduct = errors.New("no product found")

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProduct(id string) (Product, error)
	AddProduct(product Product) error
	UpdateProduct(id string, product Product) ([]Product, error)
	DeleteProduct(id string) error
}

type Store struct{}

// Returns a list of all the products
func (ps Store) GetProducts() ([]Product, error) {
	data, err := os.ReadFile(DataLocation)
	if err != nil {
		return nil, err
	}

	var allProducts []Product

	err = json.Unmarshal(data, &allProducts)
	if err != nil {
		return nil, err
	}

	return allProducts, nil
}

// Adds a new product
func (ps Store) AddProduct(product Product) error {
	productList, err := ps.GetProducts()
	if err != nil {
		return err
	}

	productList = append(productList, product)
	updatedProducts, err := json.Marshal(productList)
	if err != nil {
		return err
	}

	err = os.WriteFile(DataLocation, updatedProducts, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Gets an individual product
func (ps Store) GetProduct(id string) (Product, error) {
	productList, err := ps.GetProducts()
	if err != nil {
		return Product{}, err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			return productList[i], nil
		}
	}

	return Product{}, ErrNoProduct
}

// Updates the product with the provided id
func (ps Store) UpdateProduct(id string, product Product) ([]Product, error) {
	productList, err := ps.GetProducts()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			if productList[i].ID != product.ID {
				return nil, errors.New("updated id did not match original")
			}

			productList[i] = product

			err = overwriteProducts(productList)
			if err != nil {
				return nil, err
			}

			updatedProductList, err := ps.GetProducts()
			if err != nil {
				return nil, err
			}

			return updatedProductList, nil
		}
	}

	return nil, ErrNoProduct
}

// Removes a product; hard deletion
func (ps Store) DeleteProduct(id string) error {
	productList, err := ps.GetProducts()
	if err != nil {
		return err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			productList, err = removeElement(productList, i)
			if err != nil {
				return err
			}

			err = overwriteProducts(productList)
			if err != nil {
				return err
			}
		}
	}

	return ErrNoProduct
}

// Overwrites the entire products file; used for updating or deleting
func overwriteProducts(products []Product) error {
	updatedProducts, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = os.WriteFile(DataLocation, updatedProducts, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Removes one element from a list of products; used for deletion
func removeElement(productList []Product, index int) ([]Product, error) {
	if index < 0 || index >= len(productList) {
		return productList, errors.New("invalid index for deletion")
	}

	var updatedProducts []Product = make([]Product, 0)
	updatedProducts = append(updatedProducts, productList[0:index]...)
	updatedProducts = append(updatedProducts, productList[index+1:]...)

	return updatedProducts, nil
}
