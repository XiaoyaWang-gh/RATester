package proxy

import (
	"fmt"
	"testing"
)

func TestStop_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &Wrapper{
		closeCh: make(chan struct{}),
	}

	w.Stop()

	select {
	case <-w.closeCh:
	default:
		t.Error("closeCh should be closed")
	}
}
