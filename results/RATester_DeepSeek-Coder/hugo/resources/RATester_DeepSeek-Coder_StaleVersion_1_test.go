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

	as := &AtomicStaler{
		stale: 123,
	}

	staleVersion := as.StaleVersion()
	if staleVersion != 123 {
		t.Errorf("Expected stale version to be 123, got %d", staleVersion)
	}
}
