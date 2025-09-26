package metrics

import (
	"fmt"
	"testing"
)

func TestNewProxy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var metrics noopServerMetrics
	var proxyName string = "proxyName"
	var serverName string = "serverName"

	// Act
	metrics.NewProxy(proxyName, serverName)

	// Assert
	// TODO
}
