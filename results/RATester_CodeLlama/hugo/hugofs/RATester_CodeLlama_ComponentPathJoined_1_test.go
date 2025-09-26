package hugofs

import (
	"fmt"
	"testing"
)

func TestComponentPathJoined_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ComponentPath{
		Component: "component",
		Path:      "path",
	}
	expected := "component/path"
	if c.ComponentPathJoined() != expected {
		t.Errorf("expected %q, got %q", expected, c.ComponentPathJoined())
	}
}
