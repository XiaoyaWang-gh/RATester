package strings

import (
	"fmt"
	"testing"
)

func TestReplace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	s := "hello world"
	old := "world"
	new := "earth"
	limit := 1
	expected := "hello earth"
	result, err := ns.Replace(s, old, new, limit)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
