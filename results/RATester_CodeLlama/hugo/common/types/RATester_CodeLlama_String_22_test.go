package types

import (
	"fmt"
	"testing"
)

func TestString_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	k := KeyValues{
		Key:    "key",
		Values: []any{"value1", "value2"},
	}
	expected := "key: [value1 value2]"
	if k.String() != expected {
		t.Errorf("expected %q but got %q", expected, k.String())
	}
}
