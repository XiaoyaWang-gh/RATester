package maps

import (
	"fmt"
	"testing"
)

func TestGetMergeStrategy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Params{
		MergeStrategyKey: ParamsMergeStrategyDeep,
	}
	s, found := p.GetMergeStrategy()
	if !found {
		t.Fatal("expected to find merge strategy")
	}
	if s != ParamsMergeStrategyDeep {
		t.Fatalf("expected %q, got %q", ParamsMergeStrategyDeep, s)
	}
}
