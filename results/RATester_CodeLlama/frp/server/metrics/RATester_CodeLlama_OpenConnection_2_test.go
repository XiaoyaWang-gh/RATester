package metrics

import (
	"fmt"
	"testing"
)

func TestOpenConnection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var metrics noopServerMetrics
	var server string
	var client string

	// Act
	metrics.OpenConnection(server, client)

	// Assert
	// TODO
}
