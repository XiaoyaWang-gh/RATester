package math

import (
	"fmt"
	"testing"
)

func TestCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	if ns.Counter() != 1 {
		t.Errorf("Counter() = %d, want 1", ns.Counter())
	}
	if ns.Counter() != 2 {
		t.Errorf("Counter() = %d, want 2", ns.Counter())
	}
}
