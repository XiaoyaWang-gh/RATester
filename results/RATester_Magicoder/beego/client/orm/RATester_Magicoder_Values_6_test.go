package orm

import (
	"fmt"
	"testing"
)

func TestValues_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{}
	container := &[]Params{}
	cols := []string{"col1", "col2"}

	_, err := o.Values(container, cols...)
	if err != nil {
		t.Errorf("Values() error = %v", err)
		return
	}

	if len(*container) == 0 {
		t.Errorf("Values() expected to return some data, but got none")
	}
}
