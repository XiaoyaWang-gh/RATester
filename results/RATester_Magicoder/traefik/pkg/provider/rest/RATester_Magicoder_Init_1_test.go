package rest

import (
	"fmt"
	"testing"
)

func TestInit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Insecure: true,
	}

	err := provider.Init()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
