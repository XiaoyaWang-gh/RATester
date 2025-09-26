package metrics

import (
	"fmt"
	"testing"
)

func TestCloseProxy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var metrics noopServerMetrics
	var proxyName string = "proxyName"
	var clientName string = "clientName"

	// Act
	metrics.CloseProxy(proxyName, clientName)

	// Assert
	// TODO
}
