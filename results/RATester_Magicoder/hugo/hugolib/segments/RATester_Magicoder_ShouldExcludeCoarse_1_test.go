package segments

import (
	"fmt"
	"testing"
)

func TestShouldExcludeCoarse_1(t *testing.T) {
	tests := []struct {
		name   string
		filter segmentFilter
		field  SegmentMatcherFields
		want   bool
	}{
		{
			name: "Exclude when coarse predicate returns true",
			filter: segmentFilter{
				coarse: func(field SegmentMatcherFields) bool {
					return true
				},
			},
			field: SegmentMatcherFields{
				Kind: "test",
				Path: "test",
				Lang: "test",
			},
			want: true,
		},
		{
			name: "Include when coarse predicate returns false",
			filter: segmentFilter{
				coarse: func(field SegmentMatcherFields) bool {
					return false
				},
			},
			field: SegmentMatcherFields{
				Kind: "test",
				Path: "test",
				Lang: "test",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.filter.ShouldExcludeCoarse(tt.field); got != tt.want {
				t.Errorf("ShouldExcludeCoarse() = %v, want %v", got, tt.want)
			}
		})
	}
}
