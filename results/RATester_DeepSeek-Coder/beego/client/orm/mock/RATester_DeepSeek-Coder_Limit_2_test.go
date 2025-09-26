package mock

import (
	"fmt"
	"testing"
)

func TestLimit_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	limit := 10
	offset := 0
	qs := d.Limit(limit, offset)
	if qs != d {
		t.Errorf("Expected %v, got %v", d, qs)
	}
}
