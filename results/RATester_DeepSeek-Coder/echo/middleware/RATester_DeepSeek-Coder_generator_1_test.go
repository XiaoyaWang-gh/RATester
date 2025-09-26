package middleware

import (
	"fmt"
	"testing"
)

func TestGenerator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := generator()
	if len(result) != 32 {
		t.Errorf("Expected length of 32, got %d", len(result))
	}
}
