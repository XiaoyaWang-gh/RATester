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
		t.Errorf("Close() error = %v, wantErr %v", err, nil)
		return
	}

	if mux.isClose != true {
		t.Errorf("Close() isClose = %v, want %v", mux.isClose, true)
	}
}
