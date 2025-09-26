package migration

import (
	"fmt"
	"testing"
)

func TestSwap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := dataSlice{
		data{1, "a", nil},
		data{2, "b", nil},
		data{3, "c", nil},
	}
	d.Swap(0, 1)
	if d[0].name != "b" {
		t.Errorf("Expected d[0].name to be 'b', got %s", d[0].name)
	}
	if d[1].name != "a" {
		t.Errorf("Expected d[1].name to be 'a', got %s", d[1].name)
	}
}
