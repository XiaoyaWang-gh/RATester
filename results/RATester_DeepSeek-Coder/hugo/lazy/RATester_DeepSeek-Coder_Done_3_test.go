package lazy

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestDone_3(t *testing.T) {
	o := &onceMore{}

	t.Run("TestDone_NotDone", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o.mu.Lock()
		defer o.mu.Unlock()
		atomic.StoreUint32(&o.done, 0)
		if o.Done() {
			t.Errorf("Expected Done() to return false, got true")
		}
	})

	t.Run("TestDone_Done", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o.mu.Lock()
		defer o.mu.Unlock()
		atomic.StoreUint32(&o.done, 1)
		if !o.Done() {
			t.Errorf("Expected Done() to return true, got false")
		}
	})
}
