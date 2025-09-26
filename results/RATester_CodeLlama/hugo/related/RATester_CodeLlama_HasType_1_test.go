package related

import (
	"fmt"
	"testing"
)

func TestHasType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Config{
		Indices: IndicesConfig{
			{Type: "foo"},
			{Type: "bar"},
		},
	}

	if !c.HasType("foo") {
		t.Error("expected true")
	}

	if c.HasType("baz") {
		t.Error("expected false")
	}
}
