package mock

import (
	"fmt"
	"testing"
)

func TestOffset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	offset := 10
	qs := d.Offset(offset)
	if qs != d {
		t.Errorf("Expected %v, got %v", d, qs)
	}
}
