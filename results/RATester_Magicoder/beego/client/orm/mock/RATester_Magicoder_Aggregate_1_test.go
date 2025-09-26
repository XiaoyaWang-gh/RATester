package mock

import (
	"fmt"
	"testing"
)

func TestAggregate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	result := d.Aggregate("dept_name,sum(salary) as total")
	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}
}
