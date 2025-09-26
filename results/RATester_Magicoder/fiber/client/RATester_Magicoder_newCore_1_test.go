package client

import (
	"fmt"
	"testing"
)

func TestnewCore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	core := newCore()

	if core == nil {
		t.Error("Expected a non-nil core, but got nil")
	}

	if core.client != nil {
		t.Error("Expected client to be nil, but got a non-nil value")
	}

	if core.req != nil {
		t.Error("Expected req to be nil, but got a non-nil value")
	}

	if core.ctx != nil {
		t.Error("Expected ctx to be nil, but got a non-nil value")
	}
}
