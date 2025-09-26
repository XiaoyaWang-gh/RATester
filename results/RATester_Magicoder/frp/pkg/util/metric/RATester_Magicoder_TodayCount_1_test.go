package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestTodayCount_1(t *testing.T) {
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

	counter.Inc(1)
	counter.Inc(1)
	counter.Inc(1)

	count := counter.TodayCount()
	if count != 3 {
		t.Errorf("Expected count to be 3, but got %d", count)
	}
}
