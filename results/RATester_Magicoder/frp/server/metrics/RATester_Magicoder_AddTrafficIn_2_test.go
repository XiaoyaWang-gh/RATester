package metrics

import (
	"fmt"
	"testing"
)

func TestAddTrafficIn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	noop := noopServerMetrics{}
	noop.AddTrafficIn("", "", 0)
}
