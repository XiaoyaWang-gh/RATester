package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestNewStandardDateCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reserveDays := int64(7)
	counter := newStandardDateCounter(reserveDays)

	if counter.reserveDays != reserveDays {
		t.Errorf("Expected reserveDays to be %d, but got %d", reserveDays, counter.reserveDays)
	}

	if len(counter.counts) != int(reserveDays) {
		t.Errorf("Expected counts to have length %d, but got %d", reserveDays, len(counter.counts))
	}

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if !counter.lastUpdateDate.Equal(now) {
		t.Errorf("Expected lastUpdateDate to be %v, but got %v", now, counter.lastUpdateDate)
	}
}
