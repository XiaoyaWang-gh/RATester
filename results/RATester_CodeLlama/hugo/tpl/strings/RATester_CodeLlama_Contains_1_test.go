package strings

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

	ns := &Namespace{}
	s := "hello"
	substr := "ell"
	if contains, err := ns.Contains(s, substr); err != nil {
		t.Errorf("unexpected error: %s", err)
	} else if !contains {
		t.Errorf("expected %s to contain %s", s, substr)
	}
}
