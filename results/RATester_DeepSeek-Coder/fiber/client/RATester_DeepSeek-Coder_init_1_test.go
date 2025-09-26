package client

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

	expected := New()
	if defaultClient != expected {
		t.Errorf("Expected %v, got %v", expected, defaultClient)
	}
}
