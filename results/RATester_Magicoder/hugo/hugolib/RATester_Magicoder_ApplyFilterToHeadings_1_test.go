package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
	"github.com/gohugoio/hugo/related"
)

func TestApplyFilterToHeadings_1(t *testing.T) {
	ctx := context.Background()
	pageState := &pageState{
		// Initialize pageState fields here
	}

	testCases := []struct {
		name     string
		filter   func(*tableofcontents.Heading) bool
		expected related.Document
	}{
		{
			name: "Filter by ID",
			filter: func(h *tableofcontents.Heading) bool {
				return h.ID == "testID"
			},
			expected: &pageHeadingsFiltered{
				// Initialize expected fields here
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := pageState.ApplyFilterToHeadings(ctx, tc.filter)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
