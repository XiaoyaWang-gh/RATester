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

	now := time.Now()
	reserveDays := int64(7)
	counter := &StandardDateCounter{
		reserveDays:    reserveDays,
		counts:         make([]int64, reserveDays),
		lastUpdateDate: now,
	}

	counter.Inc(10)
	snapshot := counter.Snapshot()

	if snapshot.TodayCount() != 0 {
		t.Errorf("Expected TodayCount to be 0, got %d", snapshot.TodayCount())
	}

	if snapshot.GetLastDaysCount(1)[0] != 10 {
		t.Errorf("Expected GetLastDaysCount to be 10, got %d", snapshot.GetLastDaysCount(1)[0])
	}

	snapshot.Dec(5)
	if snapshot.GetLastDaysCount(1)[0] != 5 {
		t.Errorf("Expected GetLastDaysCount to be 5, got %d", snapshot.GetLastDaysCount(1)[0])
	}

	snapshot.Clear()
	if snapshot.TodayCount() != 0 {
		t.Errorf("Expected TodayCount to be 0 after Clear, got %d", snapshot.TodayCount())
	}
}
