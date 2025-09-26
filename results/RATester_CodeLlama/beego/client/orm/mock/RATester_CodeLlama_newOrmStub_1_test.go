package mock

import (
	"fmt"
	"testing"
)

func TestNewOrmStub_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newOrmStub()
	// Assert
	if result == nil {
		t.Errorf("newOrmStub() = nil")
	}
}
