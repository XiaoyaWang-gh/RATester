package msg

import (
	"fmt"
	"testing"
)

func TestDone_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatcher{
		doneCh: make(chan struct{}),
	}

	done := d.Done()
	if done == nil {
		t.Errorf("Expected non-nil chan struct{}, got nil")
	}

	select {
	case <-done:
		t.Errorf("Expected nothing to be sent on doneCh, got a read")
	default:
	}

	close(d.doneCh)

	select {
	case <-done:
	default:
		t.Errorf("Expected to receive on doneCh, got nothing")
	}
}
