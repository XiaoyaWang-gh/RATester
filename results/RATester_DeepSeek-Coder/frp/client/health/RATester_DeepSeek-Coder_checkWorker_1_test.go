package health

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	monitor := &Monitor{
		ctx:            ctx,
		cancel:         cancel,
		checkType:      "tcp",
		interval:       1 * time.Second,
		timeout:        100 * time.Millisecond,
		maxFailedTimes: 3,
		addr:           "localhost:8080",
	}

	go monitor.checkWorker()

	time.Sleep(5 * time.Second)

	monitor.Stop()
}
