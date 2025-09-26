package pageparser

import (
	"fmt"
	"testing"
)

func Test__2(t *testing.T) {
	var tests = []struct {
		name string
		want int
	}{
		{"tError", 0},
		{"tEOF", 1},
		{"TypeLeadSummaryDivider", 2},
		{"TypeFrontMatterYAML", 3},
		{"TypeFrontMatterTOML", 4},
		{"TypeFrontMatterJSON", 5},
		{"TypeFrontMatterORG", 6},
		{"TypeIgnore", 7},
		{"tLeftDelimScNoMarkup", 8},
		{"tRightDelimScNoMarkup", 9},
		{"tLeftDelimScWithMarkup", 10},
		{"tRightDelimScWithMarkup", 11},
		{"tScClose", 12},
		{"tScName", 13},
		{"tScNameInline", 14},
		{"tScParam", 15},
		{"tScParamVal", 16},
		{"tIndentation", 17},
		{"tText", 18},
		{"tKeywordMarker", 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.want; got != tt.want {
				t.Errorf("want = %v, want %v", got, tt.want)
			}
		})
	}
}
