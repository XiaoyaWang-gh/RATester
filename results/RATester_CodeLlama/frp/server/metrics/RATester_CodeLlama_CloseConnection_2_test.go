package metrics

import (
	"fmt"
	"testing"
)

func TestCloseConnection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var metrics noopServerMetrics
	var connectionID string = "connectionID"
	var clientID string = "clientID"

	// Act
	metrics.CloseConnection(connectionID, clientID)

	// Assert
	// TODO
}
