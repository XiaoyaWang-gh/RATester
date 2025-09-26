package lazy

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestDone_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &onceMore{}

	// Test when done is not set
	if o.Done() {
		t.Error("Done() should return false when done is not set")
	}

	// Test when done is set
	atomic.StoreUint32(&o.done, 1)
	if !o.Done() {
		t.Error("Done() should return true when done is set")
	}
}
