package mock

import (
	"fmt"
	"testing"
)

func TestSetArgs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	i := []interface{}{1, "test"}
	expected := d
	result := d.SetArgs(i...)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
