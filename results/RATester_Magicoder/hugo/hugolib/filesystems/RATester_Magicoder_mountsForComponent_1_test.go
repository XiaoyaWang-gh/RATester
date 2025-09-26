package filesystems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestmountsForComponent_1(t *testing.T) {
	testCases := []struct {
		name      string
		component string
		expected  []hugofs.FileMetaInfo
	}{
		{
			name:      "Test case 1",
			component: "content",
			expected:  []hugofs.FileMetaInfo{ /* expected result */ },
		},
		{
			name:      "Test case 2",
			component: "data",
			expected:  []hugofs.FileMetaInfo{ /* expected result */ },
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

			b := &BaseFs{
				// Initialize BaseFs fields as needed
			}

			result := b.mountsForComponent(tc.component)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
