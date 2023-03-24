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

func GetProducts() ([]byte, error) {
	data, err := os.ReadFile(DataLocation)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func AddProduct(product Product) error {
	data, err := GetProducts()
	if err != nil {
		return err
	}

	var productList []Product

	err = json.Unmarshal(data, &productList)
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

func GetProduct(id string) (Product, error) {
	data, err := GetProducts()
	if err != nil {
		return Product{}, err
	}

	var productList []Product
	err = json.Unmarshal(data, &productList)
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

func UpdateProduct(id string, product Product) ([]byte, error) {
	data, err := GetProducts()
	if err != nil {
		return []byte{}, err
	}

	var productList []Product
	err = json.Unmarshal(data, &productList)
	if err != nil {
		return []byte{}, err
	}

	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			if productList[i].ID != product.ID {
				return []byte{}, errors.New("updated id did not match original")
			}

			productList[i] = product

			err = OverwriteProducts(productList)
			if err != nil {
				return []byte{}, err
			}

			data, err := GetProducts()
			if err != nil {
				return []byte{}, err
			}

			return data, nil
		}
	}

	return []byte{}, ErrNoProduct
}

func DeleteProduct(id string) error {
	data, err := GetProducts()
	if err != nil {
		return err
	}

	var productList []Product
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

			err = OverwriteProducts(productList)
			if err != nil {
				return err
			}
		}
	}

	return ErrNoProduct
}

func removeElement(productList []Product, index int) ([]Product, error) {
	if index < 0 || index >= len(productList) {
		return productList, errors.New("invalid index for deletion")
	}

	var updatedProducts []Product = make([]Product, 0)
	updatedProducts = append(updatedProducts, productList[0:index]...)
	updatedProducts = append(updatedProducts, productList[index+1:]...)

	return updatedProducts, nil
}

func OverwriteProducts(products []Product) error {
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
