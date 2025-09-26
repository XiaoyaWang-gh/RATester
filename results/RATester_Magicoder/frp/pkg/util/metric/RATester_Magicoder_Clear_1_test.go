package metric

import (
	"fmt"
	"testing"
)

func TestClear_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &StandardDateCounter{
		reserveDays: 7,
		counts:      make([]int64, 7),
	}

	// Initialize counts
	for i := 0; i < int(counter.reserveDays); i++ {
		counter.counts[i] = int64(i + 1)
	}

	counter.Clear()

	for i := 0; i < int(counter.reserveDays); i++ {
		if counter.counts[i] != 0 {
			t.Errorf("Expected count at index %d to be 0, but got %d", i, counter.counts[i])
		}
	}
}
