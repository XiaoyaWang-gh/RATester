package resources

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
)

func TestProcess_1(t *testing.T) {
	testCases := []struct {
		name     string
		spec     string
		expected images.ImageResource
		err      error
	}{
		{
			name:     "valid spec",
			spec:     "crop:100x100",
			expected: &imageResource{},
			err:      nil,
		},
		{
			name:     "invalid spec",
			spec:     "invalid",
			expected: nil,
			err:      errors.New("invalid spec"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &imageResource{}
			result, err := i.Process(tc.spec)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
		})
	}
}
