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
		t.Errorf("Expected no error, got %v", err)
	}
	if r.closed != true {
		t.Errorf("Expected closed to be true, got %v", r.closed)
	}

	// Test that closing again doesn't panic
	err = r.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if r.closed != true {
		t.Errorf("Expected closed to be true, got %v", r.closed)
	}
}
