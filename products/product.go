package products

import (
	"encoding/json"
	"os"
)

const DataLocation = "./data/data.json"

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
