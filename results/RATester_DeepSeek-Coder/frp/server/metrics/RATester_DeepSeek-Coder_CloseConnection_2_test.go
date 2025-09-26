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

	noop := noopServerMetrics{}

	// Test that the function does not panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseConnection() panicked: %v", r)
			}
		}()

		noop.CloseConnection("", "")
	}()
}
