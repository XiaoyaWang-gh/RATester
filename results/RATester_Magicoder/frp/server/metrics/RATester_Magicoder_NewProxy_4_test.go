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

	noop := noopServerMetrics{}
	noop.NewProxy("", "")
}
