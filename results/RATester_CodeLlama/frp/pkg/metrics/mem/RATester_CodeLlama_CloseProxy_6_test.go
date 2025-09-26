package mem

import (
	"fmt"
	"testing"
)

func TestCloseProxy_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &serverMetrics{}
	m.CloseProxy("test", "test")
}
