package web

import (
	"fmt"
	"testing"
)

func TestMapping_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		methodMapping: make(map[string]func()),
	}

	testFunc := func() {
		fmt.Println("Test function")
	}

	ctrl.Mapping("GET", testFunc)

	if _, ok := ctrl.methodMapping["GET"]; !ok {
		t.Errorf("Expected method 'GET' to be mapped")
	}
}
