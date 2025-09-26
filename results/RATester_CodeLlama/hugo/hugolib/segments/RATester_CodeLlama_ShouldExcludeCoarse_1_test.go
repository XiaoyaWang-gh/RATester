package segments

import (
	"fmt"
	"testing"
)

func TestShouldExcludeCoarse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := segmentFilter{
		coarse: func(field SegmentMatcherFields) bool {
			return field.Kind == "coarse"
		},
	}

	if f.ShouldExcludeCoarse(SegmentMatcherFields{Kind: "coarse"}) {
		t.Error("ShouldExcludeCoarse should return true for a coarse field")
	}

	if !f.ShouldExcludeCoarse(SegmentMatcherFields{Kind: "fine"}) {
		t.Error("ShouldExcludeCoarse should return false for a fine field")
	}
}
