package mock

import (
	"fmt"
	"testing"
)

func TestFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}

	// Test case 1: Filter with valid input
	expected := d
	actual := d.Filter("UserName", "slene")
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case 2: Filter with invalid input
	expected = d
	actual = d.Filter("UserName", 123)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
