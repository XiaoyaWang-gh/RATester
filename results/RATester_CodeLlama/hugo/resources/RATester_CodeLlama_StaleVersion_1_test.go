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

	s := &AtomicStaler{}
	s.MarkStale()
	if s.StaleVersion() != 1 {
		t.Errorf("StaleVersion() = %d, want 1", s.StaleVersion())
	}
}
