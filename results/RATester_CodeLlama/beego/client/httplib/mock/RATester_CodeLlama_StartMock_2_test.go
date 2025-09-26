package mock

import (
	"fmt"
	"testing"
)

func TestStartMock_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var mockFilter Stub
	// Act
	stub := StartMock()
	// Assert
	if stub != mockFilter {
		t.Errorf("StartMock() = %v, want %v", stub, mockFilter)
	}
}
