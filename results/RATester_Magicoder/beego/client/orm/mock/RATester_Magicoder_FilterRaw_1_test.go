package mock

import (
	"fmt"
	"testing"
)

func TestFilterRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}

	expected := d
	actual := d.FilterRaw("s", "s2")

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
