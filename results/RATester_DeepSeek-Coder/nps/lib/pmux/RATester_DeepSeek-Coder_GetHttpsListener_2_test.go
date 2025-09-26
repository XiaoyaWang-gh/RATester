package pmux

import (
	"fmt"
	"testing"
)

func TestGetHttpsListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mux := &PortMux{
		httpsConn: make(chan *PortConn),
	}

	listener := mux.GetHttpsListener()

	if listener == nil {
		t.Errorf("Expected a non-nil listener, got nil")
	}
}
