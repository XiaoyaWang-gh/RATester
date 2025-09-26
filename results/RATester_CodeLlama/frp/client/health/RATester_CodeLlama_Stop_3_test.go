package health

import (
	"fmt"
	"testing"
)

func TestStop_3(t *testing.T) {
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
	if monitor.ctx != nil {
		t.Errorf("monitor.ctx should be nil")
	}
}
