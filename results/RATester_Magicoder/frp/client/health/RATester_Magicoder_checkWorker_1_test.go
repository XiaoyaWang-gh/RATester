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
	defer cancel()

	monitor := &Monitor{
		checkType:      "http",
		interval:       1 * time.Second,
		timeout:        1 * time.Second,
		maxFailedTimes: 3,
		url:            "http://example.com",
		ctx:            ctx,
		cancel:         cancel,
	}

	go monitor.checkWorker()

	time.Sleep(5 * time.Second)
	monitor.cancel()
}
