package pmux

import (
	"fmt"
	"testing"
)

func TestClose_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &PortListener{isClose: false}

	err := listener.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if !listener.isClose {
		t.Errorf("Expected isClose to be true, got false")
	}

	err = listener.Close()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "the listener has closed" {
		t.Errorf("Expected error 'the listener has closed', got '%v'", err)
	}
	if listener.isClose != true {
		t.Errorf("Expected isClose to be true, got false")
	}
}
