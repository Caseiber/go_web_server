package test

import (
	"go_web_server/stub"
	"testing"
)

var store = stub.StubStore{}

func checkResponseCode(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Wrong status code returned, expected %d, received %d", expected, actual)
	}
}
