package msg

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewDispatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &bytes.Buffer{}
	d := NewDispatcher(rw)
	if d.rw != rw {
		t.Errorf("NewDispatcher() rw = %v, want %v", d.rw, rw)
	}
	if d.sendCh == nil {
		t.Errorf("NewDispatcher() sendCh = nil, want not nil")
	}
	if d.doneCh == nil {
		t.Errorf("NewDispatcher() doneCh = nil, want not nil")
	}
	if d.msgHandlers == nil {
		t.Errorf("NewDispatcher() msgHandlers = nil, want not nil")
	}
}
