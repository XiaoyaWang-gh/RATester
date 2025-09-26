package helpers

import (
	"fmt"
	"testing"
)

func TestAdd_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stats := &ProcessingStats{}
	counter := uint64(10)
	amount := 5
	stats.Add(&counter, amount)
	if counter != 15 {
		t.Errorf("Expected counter to be 15, but got %d", counter)
	}
}
