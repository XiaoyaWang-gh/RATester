package mock

import (
	"fmt"
	"testing"
)

func TestOrderBy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}

	expected := d
	actual := d.OrderBy("id")

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
