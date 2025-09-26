package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestRotate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	c := &StandardDateCounter{
		reserveDays:    10,
		counts:         make([]int64, 10),
		lastUpdateDate: now,
	}
	c.rotate(now)
	if len(c.counts) != 10 {
		t.Errorf("rotate failed, len(c.counts) = %d", len(c.counts))
	}
}
