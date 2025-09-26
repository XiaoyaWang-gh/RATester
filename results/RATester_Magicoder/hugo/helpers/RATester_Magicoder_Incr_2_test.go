package helpers

import (
	"fmt"
	"testing"
)

func TestIncr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stats := ProcessingStats{}
	var counter uint64
	stats.Incr(&counter)
	if counter != 1 {
		t.Errorf("Expected counter to be 1, but got %d", counter)
	}
}
