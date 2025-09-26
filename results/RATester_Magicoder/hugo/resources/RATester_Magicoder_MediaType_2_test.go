package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestMediaType_2(t *testing.T) {
	type testCase struct {
		name     string
		resource *genericResource
		expected media.Type
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			resource: &genericResource{
				sd: ResourceSourceDescriptor{
					MediaType: media.Type{
						Type: "application/rss+xml",
					},
				},
			},
			expected: media.Type{
				Type: "application/rss+xml",
			},
		},
		{
			name: "Test case 2",
			resource: &genericResource{
				sd: ResourceSourceDescriptor{
					MediaType: media.Type{
						Type: "text/html",
					},
				},
			},
			expected: media.Type{
				Type: "text/html",
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

			result := tc.resource.MediaType()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
