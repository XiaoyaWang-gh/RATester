package lazy

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestInProgress_1(t *testing.T) {
	om := &onceMore{}

	t.Run("InProgress should return false when lock is not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if om.InProgress() {
			t.Errorf("Expected InProgress() to return false, got true")
		}
	})

	t.Run("InProgress should return true when lock is set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		atomic.StoreUint32(&om.lock, 1)
		if !om.InProgress() {
			t.Errorf("Expected InProgress() to return true, got false")
		}
	})
}
