package client

import (
	"fmt"
	"testing"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctl := &Control{
		doneCh: make(chan struct{}),
	}

	doneCh := ctl.Done()

	select {
	case <-doneCh:
		t.Errorf("Expected the channel to be open")
	default:
		// The channel is open, so we pass the test
	}

	close(ctl.doneCh)

	select {
	case <-doneCh:
		// The channel is closed, so we pass the test
	default:
		t.Errorf("Expected the channel to be closed")
	}
}
