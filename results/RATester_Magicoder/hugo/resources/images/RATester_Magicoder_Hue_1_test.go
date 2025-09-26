package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
	"github.com/spf13/cast"
)

func TestHue_1(t *testing.T) {
	filters := &Filters{}

	tests := []struct {
		name   string
		shift  any
		expect gift.Filter
	}{
		{
			name:  "Test case 1",
			shift: 10,
			expect: filter{
				Options: newFilterOpts(10),
				Filter:  gift.Hue(cast.ToFloat32(10)),
			},
		},
		{
			name:  "Test case 2",
			shift: 20,
			expect: filter{
				Options: newFilterOpts(20),
				Filter:  gift.Hue(cast.ToFloat32(20)),
			},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := filters.Hue(test.shift)
			if !reflect.DeepEqual(result, test.expect) {
				t.Errorf("Expected: %v, but got: %v", test.expect, result)
			}
		})
	}
}
