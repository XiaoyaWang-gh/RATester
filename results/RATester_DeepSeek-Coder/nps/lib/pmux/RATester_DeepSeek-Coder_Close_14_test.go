package pmux

import (
	"fmt"
	"testing"
)

func TestClose_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mux := &PortMux{
		isClose: false,
	}

	err := mux.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if mux.isClose != true {
		t.Errorf("Expected isClose to be true, got %v", mux.isClose)
	}

	mux.isClose = true
	err = mux.Close()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if mux.isClose != true {
		t.Errorf("Expected isClose to be true, got %v", mux.isClose)
	}
}
