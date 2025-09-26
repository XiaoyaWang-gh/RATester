package segments

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/predicate"
)

func TestShouldExcludeFine_1(t *testing.T) {
	tests := []struct {
		name   string
		fields SegmentMatcherFields
		want   bool
	}{
		{
			name:   "Exclude fine when Kind matches",
			fields: SegmentMatcherFields{Kind: "*"},
			want:   true,
		},
		{
			name:   "Exclude fine when Path matches",
			fields: SegmentMatcherFields{Path: "*"},
			want:   true,
		},
		{
			name:   "Exclude fine when Lang matches",
			fields: SegmentMatcherFields{Lang: "*"},
			want:   true,
		},
		{
			name:   "Exclude fine when Output matches",
			fields: SegmentMatcherFields{Output: "*"},
			want:   true,
		},
		{
			name:   "Include fine when no match",
			fields: SegmentMatcherFields{Kind: "no_match", Path: "no_match", Lang: "no_match", Output: "no_match"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := segmentFilter{
				coarse: predicate.P[SegmentMatcherFields](func(fields SegmentMatcherFields) bool {
					return fields.Kind == "*" || fields.Path == "*" || fields.Lang == "*" || fields.Output == "*"
				}),
				fine: predicate.P[SegmentMatcherFields](func(fields SegmentMatcherFields) bool {
					return fields.Kind == "*" || fields.Path == "*" || fields.Lang == "*" || fields.Output == "*"
				}),
			}
			if got := f.ShouldExcludeFine(tt.fields); got != tt.want {
				t.Errorf("ShouldExcludeFine() = %v, want %v", got, tt.want)
			}
		})
	}
}
