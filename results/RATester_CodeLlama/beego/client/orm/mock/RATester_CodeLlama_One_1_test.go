package mock

import (
	"fmt"
	"testing"
)

func TestOne_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	container := &struct{}{}
	cols := []string{"col1", "col2"}
	err := d.One(container, cols...)
	if err != nil {
		t.Errorf("One() error = %v", err)
		return
	}
}
