package aggregate

import (
	"fmt"
	"testing"
)

func TestNewProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &serverMetrics{}
	m.NewProxy("name", "proxyType")
}
