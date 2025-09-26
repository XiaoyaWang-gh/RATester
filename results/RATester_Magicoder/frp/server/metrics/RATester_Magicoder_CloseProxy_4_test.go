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
	noop.CloseProxy("", "")
}
