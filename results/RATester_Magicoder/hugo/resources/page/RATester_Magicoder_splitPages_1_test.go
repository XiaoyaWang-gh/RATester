package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsplitPages_1(t *testing.T) {
	pages := []Page{
		// create some pages
	}

	testCases := []struct {
		name     string
		pages    Pages
		size     int
		expected []paginatedElement
	}{
		{
			name:     "Test case 1",
			pages:    pages,
			size:     10,
			expected: []paginatedElement{
				// create expected paginated elements
			},
		},
		{
			name:     "Test case 2",
			pages:    pages,
			size:     20,
			expected: []paginatedElement{
				// create expected paginated elements
			},
		},
		// add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := splitPages(tc.pages, tc.size)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
