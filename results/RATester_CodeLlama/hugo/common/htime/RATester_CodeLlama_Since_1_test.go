package htime

import (
	"fmt"
	"testing"
	"time"
)

func TestSince_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	t0 := time.Now()
	d := Since(t0)
	if d < 0 {
		t.Errorf("Since(%v) = %v; want non-negative", t0, d)
	}
}
