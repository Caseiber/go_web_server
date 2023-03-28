package stub

type ErrorValue int64

// Defines constants to be used for unit tests
const (
	Success ErrorValue = iota
	ErrorNoProduct
)
