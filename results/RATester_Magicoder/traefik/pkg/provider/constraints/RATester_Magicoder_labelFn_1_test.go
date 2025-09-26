package constraints

import (
	"fmt"
	"testing"
)

func TestLabelFn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "testName"
	value := "testValue"
	labels := map[string]string{
		name: value,
	}
	fn := labelFn(name, value)
	if !fn(labels) {
		t.Errorf("Expected true, got false")
	}
}
