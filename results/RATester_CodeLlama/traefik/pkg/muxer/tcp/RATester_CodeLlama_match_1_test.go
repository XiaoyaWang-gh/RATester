package tcp

import (
	"fmt"
	"testing"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	// When
	// Then
}
