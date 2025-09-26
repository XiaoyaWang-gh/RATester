package strings

import (
	"fmt"
	"testing"
)

func TestHasSuffix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	s := "hello"
	suffix := "world"
	expected := true
	result, err := ns.HasSuffix(s, suffix)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
