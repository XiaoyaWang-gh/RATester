package maps

import (
	"fmt"
	"testing"
)

func TestDeleteMergeStrategy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Params{
		"a": "b",
		"c": "d",
	}
	if !p.DeleteMergeStrategy() {
		t.Error("expected true")
	}
	if p.IsZero() {
		t.Error("expected false")
	}
	if _, found := p[MergeStrategyKey]; found {
		t.Error("expected false")
	}
	if _, found := p["a"]; !found {
		t.Error("expected true")
	}
	if _, found := p["c"]; !found {
		t.Error("expected true")
	}
}
