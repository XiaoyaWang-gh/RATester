package testhelpers

import (
	"fmt"
	"testing"
)

func TestNewCollectingHealthCheckMetrics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	collectingHealthCheckMetrics := NewCollectingHealthCheckMetrics()
	// Assert
	if collectingHealthCheckMetrics == nil {
		t.Errorf("NewCollectingHealthCheckMetrics() should not return nil")
	}
}
