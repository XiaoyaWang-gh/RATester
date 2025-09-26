package metric

import (
	"fmt"
	"testing"
)

func TestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &StandardCounter{count: 10}
	result := counter.Count()
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}
