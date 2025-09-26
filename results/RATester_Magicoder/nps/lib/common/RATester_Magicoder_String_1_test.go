package common

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	addr := &Addr{
		Type: 1,
		Host: "localhost",
		Port: 8080,
	}
	expected := "localhost:8080"
	result := addr.String()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
