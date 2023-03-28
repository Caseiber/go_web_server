package stub

import "go_web_server/products"

// Stubbing out the store to allow for unit testing
type Stub struct{}

func (s Stub) GetProducts() ([]byte, error) {
	return []byte{}, nil
}

func (s Stub) GetProduct(id int) (products.Product, error) {
	if id == int(ErrorNoProduct) {
		return products.Product{}, products.ErrNoProduct
	}
	return products.Product{}, nil
}

func (s Stub) AddProduct(product products.Product) error {
	if product.ID == int(ErrorNoProduct) {
		return products.ErrNoProduct
	}

	return nil
}

func (s Stub) UpdateProduct(id int, product products.Product) ([]byte, error) {
	if id == int(ErrorNoProduct) {
		return []byte{}, products.ErrNoProduct
	}

	return []byte{}, nil
}

func (s Stub) DeleteProduct(id int) error { return nil }
