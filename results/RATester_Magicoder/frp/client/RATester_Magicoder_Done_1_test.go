package client

import (
	"fmt"
	"testing"
	"time"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{
		doneCh: make(chan struct{}),
	}

	doneCh := ctl.Done()

	go func() {
		time.Sleep(time.Second)
		close(ctl.doneCh)
	}()

	select {
	case <-doneCh:
		// Test passed
		return
	case <-time.After(2 * time.Second):
		t.Error("Test timed out")
	}
}
