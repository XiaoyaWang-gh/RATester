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
	result := d.Offset(offset)
	if result != d {
		t.Errorf("Expected %v, but got %v", d, result)
	}
}
