package gin

import (
	"fmt"
	"testing"
)

func TestGet_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := Params{
		Param{Key: "name", Value: "Alice"},
		Param{Key: "age", Value: "29"},
	}
	if v, ok := ps.Get("name"); !ok || v != "Alice" {
		t.Errorf("Get(%q) = %q, %t want %q, %t", "name", v, ok, "Alice", true)
	}
	if v, ok := ps.Get("age"); !ok || v != "29" {
		t.Errorf("Get(%q) = %q, %t want %q, %t", "age", v, ok, "29", true)
	}
	if v, ok := ps.Get("unknown"); ok || v != "" {
		t.Errorf("Get(%q) = %q, %t want %q, %t", "unknown", v, ok, "", false)
	}
}
