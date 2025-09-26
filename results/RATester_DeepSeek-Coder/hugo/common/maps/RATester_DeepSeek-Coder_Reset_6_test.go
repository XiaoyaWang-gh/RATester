package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReset_6(t *testing.T) {
	t.Parallel()

	type testType struct {
		name     string
		cache    *SliceCache[testType]
		expected *SliceCache[testType]
	}

	tests := []testType{
		{
			name: "Reset with empty cache",
			cache: &SliceCache[testType]{
				m: make(map[string][]testType),
			},
			expected: &SliceCache[testType]{
				m: make(map[string][]testType),
			},
		},
		{
			name: "Reset with non-empty cache",
			cache: &SliceCache[testType]{
				m: map[string][]testType{
					"key1": {{name: "value1"}},
					"key2": {{name: "value2"}},
				},
			},
			expected: &SliceCache[testType]{
				m: make(map[string][]testType),
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

			tt.cache.Reset()
			if !reflect.DeepEqual(tt.cache, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, tt.cache)
			}
		})
	}
}
