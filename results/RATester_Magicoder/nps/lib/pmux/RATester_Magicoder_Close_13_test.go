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
		t.Errorf("Close() returned an error: %v", err)
	}

	if !listener.isClose {
		t.Error("Close() did not set isClose to true")
	}

	err = listener.Close()
	if err == nil {
		t.Error("Close() did not return an error when called again")
	}
}
