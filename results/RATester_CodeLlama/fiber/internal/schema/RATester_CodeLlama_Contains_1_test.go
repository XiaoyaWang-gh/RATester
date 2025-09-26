package schema

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := tagOptions{"a", "b", "c"}
	if !o.Contains("b") {
		t.Errorf("Contains failed. Got %v, want true", o.Contains("b"))
	}
}
