package client

import (
	"fmt"
	"testing"
	"time"
)

func TestKeepControllerWorking_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{
		ctl: &Control{
			doneCh: make(chan struct{}),
		},
	}

	go func() {
		time.Sleep(1 * time.Second)
		close(svr.ctl.doneCh)
	}()

	svr.keepControllerWorking()

	if svr.ctl != nil {
		t.Errorf("Expected ctl to be nil after keepControllerWorking, but got %v", svr.ctl)
	}
}
