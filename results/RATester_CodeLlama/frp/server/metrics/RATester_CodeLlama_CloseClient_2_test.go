package metrics

import (
	"fmt"
	"testing"
)

func TestCloseClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	noopServerMetrics := noopServerMetrics{}

	// Act
	noopServerMetrics.CloseClient()

	// Assert
	// ...
}
