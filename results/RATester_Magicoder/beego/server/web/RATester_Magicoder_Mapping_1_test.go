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

	ctrl.Mapping("GET", func() {
		// test function
	})

	if _, ok := ctrl.methodMapping["GET"]; !ok {
		t.Error("Expected methodMapping to contain 'GET'")
	}
}
