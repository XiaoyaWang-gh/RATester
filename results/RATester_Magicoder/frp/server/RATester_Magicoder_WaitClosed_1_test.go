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

	ctl := &Control{
		doneCh: make(chan struct{}),
	}

	go func() {
		time.Sleep(time.Second)
		close(ctl.doneCh)
	}()

	ctl.WaitClosed()
}
