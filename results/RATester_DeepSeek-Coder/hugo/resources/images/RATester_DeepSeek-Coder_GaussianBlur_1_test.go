package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestGaussianBlur_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		sigma    any
		expected gift.Filter
	}{
		{
			name:  "Test GaussianBlur with float32 sigma",
			sigma: float32(2.0),
			expected: filter{
				Options: newFilterOpts(float32(2.0)),
				Filter:  gift.GaussianBlur(float32(2.0)),
			},
		},
		{
			name:  "Test GaussianBlur with int sigma",
			sigma: 2,
			expected: filter{
				Options: newFilterOpts(2),
				Filter:  gift.GaussianBlur(2),
			},
		},
		{
			name:  "Test GaussianBlur with string sigma",
			sigma: "2",
			expected: filter{
				Options: newFilterOpts("2"),
				Filter:  gift.GaussianBlur(2),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &Filters{}
			result := f.GaussianBlur(tt.sigma)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
