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

	name := "name"
	value := "value"
	labelFn := labelFn(name, value)
	labels := map[string]string{
		"name": "value",
	}
	if !labelFn(labels) {
		t.Errorf("labelFn failed")
	}
}
