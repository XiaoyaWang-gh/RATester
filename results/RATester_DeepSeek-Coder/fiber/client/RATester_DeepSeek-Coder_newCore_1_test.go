package client

import (
	"fmt"
	"testing"
)

func TestNewCore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	core := newCore()

	if core.client != nil {
		t.Errorf("Expected client to be nil, got %v", core.client)
	}

	if core.req != nil {
		t.Errorf("Expected req to be nil, got %v", core.req)
	}

	if core.ctx != nil {
		t.Errorf("Expected ctx to be nil, got %v", core.ctx)
	}
}
