package server

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitClosed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctl := &Control{
		doneCh: make(chan struct{}),
	}

	go func() {
		time.Sleep(time.Second)
		close(ctl.doneCh)
	}()

	ctl.WaitClosed()

	select {
	case <-ctl.doneCh:
		t.Error("doneCh should not be closed")
	default:
		t.Log("doneCh is not closed")
	}
}
