package health

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckWorker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	monitor := &Monitor{
		checkType:      "tcp",
		interval:       time.Second,
		timeout:        time.Second,
		maxFailedTimes: 3,
		addr:           "127.0.0.1:8080",
		failedTimes:    0,
		statusOK:       true,
		statusNormalFn: func() {},
		statusFailedFn: func() {},
	}
	monitor.checkWorker()
}
