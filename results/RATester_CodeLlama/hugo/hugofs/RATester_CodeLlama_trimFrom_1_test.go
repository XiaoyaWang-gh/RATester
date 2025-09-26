package hugofs

import (
	"fmt"
	"testing"
)

func TestTrimFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := RootMapping{
		From: "blog",
	}
	name := "blog/foo.md"
	if r.trimFrom(name) != "foo.md" {
		t.Errorf("expected %q got %q", "foo.md", r.trimFrom(name))
	}
}
