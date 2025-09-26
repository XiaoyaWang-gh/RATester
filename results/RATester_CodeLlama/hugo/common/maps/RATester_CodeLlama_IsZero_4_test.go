package maps

import (
	"fmt"
	"testing"
)

func TestIsZero_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Params{}
	if !p.IsZero() {
		t.Errorf("IsZero() = %v, want %v", p.IsZero(), true)
	}

	p = Params{"a": 1}
	if p.IsZero() {
		t.Errorf("IsZero() = %v, want %v", p.IsZero(), false)
	}

	p = Params{"a": 1, "b": 2}
	if p.IsZero() {
		t.Errorf("IsZero() = %v, want %v", p.IsZero(), false)
	}

	p = Params{MergeStrategyKey: "merge"}
	if p.IsZero() {
		t.Errorf("IsZero() = %v, want %v", p.IsZero(), false)
	}
}
