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

	noop := noopServerMetrics{}

	// Test that calling CloseProxy does not panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseProxy() panicked: %v", r)
			}
		}()

		noop.CloseProxy("", "")
	}()
}
