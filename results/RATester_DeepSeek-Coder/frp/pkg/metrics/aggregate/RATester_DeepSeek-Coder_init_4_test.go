package aggregate

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/server/metrics"
)

func TestInit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	metrics.Register(sm)
}
