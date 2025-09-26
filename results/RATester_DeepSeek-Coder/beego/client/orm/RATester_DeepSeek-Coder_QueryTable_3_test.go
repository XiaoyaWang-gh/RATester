package orm

import (
	"fmt"
	"testing"
)

func TestQueryTable_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	qs := d.QueryTable(nil)
	if qs != nil {
		t.Errorf("Expected nil, got %v", qs)
	}
}
