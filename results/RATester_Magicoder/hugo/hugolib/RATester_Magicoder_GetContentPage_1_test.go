package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestGetContentPage_1(t *testing.T) {
	hugoSites := &HugoSites{
		// Initialize HugoSites fields here
	}

	testCases := []struct {
		name     string
		filename string
		expected page.Page
	}{
		{
			name:     "Test case 1",
			filename: "test1.md",
			expected: &pageState{
				// Initialize expected page fields here
			},
		},
		{
			name:     "Test case 2",
			filename: "test2.md",
			expected: &pageState{
				// Initialize expected page fields here
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

			result := hugoSites.GetContentPage(tc.filename)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
