package http

import (
	"fmt"
	"testing"
)

func TestAddress_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		addr: "localhost:8080",
	}

	expected := "localhost:8080"
	actual := server.Address()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
