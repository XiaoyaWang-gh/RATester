package metric

import (
	"fmt"
	"testing"
)

func TestNewCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := NewCounter()

	if counter == nil {
		t.Error("NewCounter() should not return nil")
	}

	if counter.Count() != 0 {
		t.Error("NewCounter() should return a counter with count 0")
	}
}
