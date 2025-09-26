package gin

import (
	"fmt"
	"testing"
)

func TestByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := Params{
		Param{Key: "name", Value: "Alice"},
		Param{Key: "age", Value: "29"},
	}
	va := ps.ByName("name")
	if va != "Alice" {
		t.Errorf("ByName(%q) = %q, want %q", "name", va, "Alice")
	}
}
