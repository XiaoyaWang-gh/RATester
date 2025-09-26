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

	noop := noopServerMetrics{}

	// Call the method under test
	noop.CloseClient()

	// No assertions are needed here because the method is a no-op
}
