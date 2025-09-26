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

	orm := &DoNothingOrm{}
	stats := orm.DBStats()
	if stats != nil {
		t.Errorf("Expected nil, got %v", stats)
	}
}
