package resources

import (
	"fmt"
	"testing"
)

func TestMarkStale_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stale := &AtomicStaler{}
	stale.MarkStale()

	if stale.StaleVersion() != 1 {
		t.Errorf("Expected stale version to be 1, got %d", stale.StaleVersion())
	}
}
