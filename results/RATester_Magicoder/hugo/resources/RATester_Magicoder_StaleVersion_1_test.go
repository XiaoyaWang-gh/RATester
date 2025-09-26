package resources

import (
	"fmt"
	"testing"
)

func TestStaleVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stale := &AtomicStaler{stale: 1}
	version := stale.StaleVersion()
	if version != 1 {
		t.Errorf("Expected stale version to be 1, but got %d", version)
	}
}
