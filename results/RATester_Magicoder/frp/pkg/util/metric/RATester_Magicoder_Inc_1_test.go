package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestInc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &StandardDateCounter{
		reserveDays:    7,
		counts:         make([]int64, 7),
		lastUpdateDate: time.Now(),
	}

	counter.Inc(10)

	if counter.counts[0] != 10 {
		t.Errorf("Expected count to be 10, but got %d", counter.counts[0])
	}
}
