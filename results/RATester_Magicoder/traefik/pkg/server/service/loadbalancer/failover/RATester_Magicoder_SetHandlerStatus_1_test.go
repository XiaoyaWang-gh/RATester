package failover

import (
	"context"
	"fmt"
	"testing"
)

func TestSetHandlerStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	f := &Failover{}

	// Test case 1: When the handler status is UP and the input is also UP
	f.handlerStatus = true
	f.SetHandlerStatus(ctx, true)
	if f.handlerStatus != true {
		t.Errorf("Expected handler status to be UP, but got %v", f.handlerStatus)
	}

	// Test case 2: When the handler status is DOWN and the input is also DOWN
	f.handlerStatus = false
	f.SetHandlerStatus(ctx, false)
	if f.handlerStatus != false {
		t.Errorf("Expected handler status to be DOWN, but got %v", f.handlerStatus)
	}

	// Test case 3: When the handler status is UP and the input is DOWN
	f.handlerStatus = true
	f.SetHandlerStatus(ctx, false)
	if f.handlerStatus != false {
		t.Errorf("Expected handler status to be DOWN, but got %v", f.handlerStatus)
	}

	// Test case 4: When the handler status is DOWN and the input is UP
	f.handlerStatus = false
	f.SetHandlerStatus(ctx, true)
	if f.handlerStatus != true {
		t.Errorf("Expected handler status to be UP, but got %v", f.handlerStatus)
	}
}
