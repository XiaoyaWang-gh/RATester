package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPut_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{}
	ctx := context.Background()
	key := "key"
	val := "val"
	timeout := time.Second
	err := bc.Put(ctx, key, val, timeout)
	if err != nil {
		t.Errorf("Put() error = %v, want nil", err)
		return
	}
	if bc.items == nil {
		t.Errorf("Put() bc.items = nil, want not nil")
		return
	}
	if bc.items[key] == nil {
		t.Errorf("Put() bc.items[key] = nil, want not nil")
		return
	}
	if bc.items[key].val != val {
		t.Errorf("Put() bc.items[key].val = %v, want %v", bc.items[key].val, val)
		return
	}
	if bc.items[key].createdTime.IsZero() {
		t.Errorf("Put() bc.items[key].createdTime = %v, want not zero", bc.items[key].createdTime)
		return
	}
	if bc.items[key].lifespan != timeout {
		t.Errorf("Put() bc.items[key].lifespan = %v, want %v", bc.items[key].lifespan, timeout)
		return
	}
}
