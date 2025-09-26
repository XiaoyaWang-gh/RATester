package mock

import (
	"fmt"
	"testing"
)

func TestAll_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	var container []interface{}
	count, err := d.All(container, "column1", "column2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if count != 0 {
		t.Errorf("Expected count to be 0, got %v", count)
	}
}
