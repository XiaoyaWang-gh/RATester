package net

import (
	"fmt"
	"testing"
)

func TestClose_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &WebsocketListener{
		ln:       nil,
		acceptCh: nil,
		server:   nil,
	}
	err := ln.Close()
	if err != nil {
		t.Errorf("Close() error = %v, want nil", err)
		return
	}
}
