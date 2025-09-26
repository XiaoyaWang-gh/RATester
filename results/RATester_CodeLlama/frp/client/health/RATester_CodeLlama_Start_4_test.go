package health

import (
	"fmt"
	"testing"
)

func TestStart_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	monitor := &Monitor{
		checkType: "tcp",
		addr:      "127.0.0.1:8080",
	}
	monitor.Start()
	monitor.Stop()
}
