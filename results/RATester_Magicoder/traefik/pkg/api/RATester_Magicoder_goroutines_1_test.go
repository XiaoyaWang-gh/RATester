package api

import (
	"fmt"
	"testing"
)

func Testgoroutines_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := goroutines()
	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}
}
