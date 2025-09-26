package metrics

import (
	"fmt"
	"testing"
)

func TestWrapEntryPointHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	// When
	// Then
}
