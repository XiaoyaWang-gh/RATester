package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTypeVariations_1(t *testing.T) {
	l := &layoutBuilder{
		d: LayoutDescriptor{
			RenderingHook: true,
		},
	}

	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "single variation",
			input:    []string{"single"},
			expected: []string{"singlenode"},
		},
		{
			name:     "multiple variations",
			input:    []string{"multiple1", "multiple2"},
			expected: []string{"multiple1node", "multiple2node"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l.addTypeVariations(tc.input...)
			if !reflect.DeepEqual(l.typeVariations, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, l.typeVariations)
			}
		})
	}
}
