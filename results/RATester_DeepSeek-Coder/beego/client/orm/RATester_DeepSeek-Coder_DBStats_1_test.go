package orm

import (
	"fmt"
	"testing"
)

func TestDBStats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	stats := d.DBStats()
	if stats != nil {
		t.Errorf("Expected nil, got %v", stats)
	}
}
