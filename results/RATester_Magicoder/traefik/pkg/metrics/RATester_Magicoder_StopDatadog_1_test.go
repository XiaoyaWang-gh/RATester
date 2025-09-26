package metrics

import (
	"fmt"
	"testing"
)

func TestStopDatadog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	datadogLoopCancelFunc = func() {
		// Mock the datadogLoopCancelFunc
	}

	// Act
	StopDatadog()

	// Assert
	if datadogLoopCancelFunc != nil {
		t.Errorf("Expected datadogLoopCancelFunc to be nil after calling StopDatadog, but it's not.")
	}
}
