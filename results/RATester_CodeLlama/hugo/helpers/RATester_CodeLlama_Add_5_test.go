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

	s := &ProcessingStats{}
	counter := uint64(0)
	amount := 10
	s.Add(&counter, amount)
	if counter != uint64(amount) {
		t.Errorf("Expected %d, got %d", amount, counter)
	}
}
