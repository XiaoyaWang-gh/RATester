package mock

import (
	"fmt"
	"testing"
)

func TestDistinct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	result := d.Distinct()
	if result != d {
		t.Errorf("Expected Distinct() to return the same instance of DoNothingQuerySetter, but got %v", result)
	}
}
