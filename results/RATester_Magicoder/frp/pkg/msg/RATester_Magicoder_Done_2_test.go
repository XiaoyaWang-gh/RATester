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

	doneCh := d.Done()

	if doneCh == nil {
		t.Error("Expected non-nil done channel")
	}
}
