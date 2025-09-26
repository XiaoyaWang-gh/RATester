package migration

import (
	"fmt"
	"testing"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := dataSlice{data{}, data{}, data{}}
	expected := 3
	actual := d.Len()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
