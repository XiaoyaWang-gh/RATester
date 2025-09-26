package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
	"github.com/gohugoio/hugo/common/maps"
)

func TestText_2(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		text     string
		options  []any
		expected gift.Filter
	}

	tests := []testCase{
		{
			name:    "default",
			text:    "Hello, World",
			options: []any{},
			expected: filter{
				Options: newFilterOpts("Hello, World", maps.Params{}),
				Filter: textFilter{
					text:        "Hello, World",
					color:       color.White,
					size:        20,
					x:           10,
					y:           10,
					linespacing: 2,
				},
			},
		},
		{
			name: "with color",
			text: "Hello, World",
			options: []any{
				"color", color.Black,
			},
			expected: filter{
				Options: newFilterOpts("Hello, World", maps.Params{"color": color.Black}),
				Filter: textFilter{
					text:        "Hello, World",
					color:       color.Black,
					size:        20,
					x:           10,
					y:           10,
					linespacing: 2,
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := new(Filters).Text(tc.text, tc.options...)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
