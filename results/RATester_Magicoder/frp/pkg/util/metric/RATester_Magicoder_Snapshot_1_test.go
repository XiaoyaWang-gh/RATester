package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestSnapshot_1(t *testing.T) {
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

	counter.counts[0] = 10
	counter.counts[1] = 20
	counter.counts[2] = 30
	counter.counts[3] = 40
	counter.counts[4] = 50
	counter.counts[5] = 60
	counter.counts[6] = 70

	snapshot := counter.Snapshot()

	if snapshot.TodayCount() != 0 {
		t.Errorf("Expected today count to be 0, but got %d", snapshot.TodayCount())
	}

	lastDaysCount := snapshot.GetLastDaysCount(7)
	if len(lastDaysCount) != 7 {
		t.Errorf("Expected last days count to be 7, but got %d", len(lastDaysCount))
	}

	for i := 0; i < 7; i++ {
		if lastDaysCount[i] != counter.counts[i] {
			t.Errorf("Expected count for day %d to be %d, but got %d", i, counter.counts[i], lastDaysCount[i])
		}
	}
}
