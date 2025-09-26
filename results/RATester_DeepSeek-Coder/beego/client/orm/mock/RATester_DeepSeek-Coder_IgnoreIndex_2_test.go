package mock

import (
	"fmt"
	"testing"
)

func TestIgnoreIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	qs := d.IgnoreIndex("test_index")
	if qs != d {
		t.Errorf("Expected IgnoreIndex to return the original DoNothingQuerySetter, but got %v", qs)
	}
}
