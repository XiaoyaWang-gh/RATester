package task

import (
	"fmt"
	"testing"
)

func TestAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := bounds{min: 1, max: 10, names: map[string]uint{"one": 1, "two": 2, "three": 3}}
	if all(r) != 1023 {
		t.Errorf("all(r) = %d, want 1023", all(r))
	}
}
