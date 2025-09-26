package filenotify

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestsendErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		errors: make(chan error),
		done:   make(chan struct{}),
	}

	testErr := errors.New("test error")

	go func() {
		w.sendErr(testErr)
		close(w.done)
	}()

	select {
	case err := <-w.errors:
		if err != testErr {
			t.Errorf("Expected error %v, got %v", testErr, err)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for error")
	}
}
