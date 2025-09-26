package metric

import (
	"fmt"
	"sync"
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
		mu:          sync.Mutex{},
	}

	counter.counts[0] = 1
	counter.counts[1] = 2
	counter.counts[2] = 3
	counter.counts[3] = 4
	counter.counts[4] = 5
	counter.counts[5] = 6
	counter.counts[6] = 7

	counter.Clear()

	for i := 0; i < int(counter.reserveDays); i++ {
		if counter.counts[i] != 0 {
			t.Errorf("Expected count at index %d to be 0, got %d", i, counter.counts[i])
		}
	}
}
