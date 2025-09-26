package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewAbsURLReplacer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newAbsURLReplacer()
	// Assert
	if result == nil {
		t.Errorf("newAbsURLReplacer() failed")
	}
}
