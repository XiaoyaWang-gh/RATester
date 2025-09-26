package tasks

import (
	"fmt"
	"testing"
)

func TestClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		closed: false,
		quit:   make(chan struct{}),
	}

	err := r.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !r.closed {
		t.Errorf("Expected closed to be true, but it was false")
	}

	if r.quit == nil {
		t.Errorf("Expected quit to be a non-nil channel, but it was nil")
	}

	err = r.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.closed {
		t.Errorf("Expected closed to be false, but it was true")
	}

	if r.quit != nil {
		t.Errorf("Expected quit to be a nil channel, but it was not")
	}
}
