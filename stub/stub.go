package stub

import (
	"fmt"
	"go_web_server/products"
)

type ErrorValue int64

var ValidProduct = products.Product{
	ID:          fmt.Sprint(Success),
	Name:        "Test Coffee",
	Description: "A coffee for to test",
	Price:       5,
	IsAvailable: true,
}

// Defines constants to be used for unit tests
const (
	Success ErrorValue = iota
	ErrorNoProduct
	ErrorBadProduct
)
