package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMapByPrefix_1(t *testing.T) {
	testSet := map[string]string{
		"apple":  "fruit",
		"banana": "fruit",
		"cat":    "animal",
		"dog":    "animal",
	}

	testCases := []struct {
		name     string
		set      map[string]string
		prefix   string
		expected map[string]string
	}{
		{
			name:     "Test case 1: Prefix 'a'",
			set:      testSet,
			prefix:   "a",
			expected: map[string]string{"apple": "fruit", "banana": "fruit"},
		},
		{
			name:     "Test case 2: Prefix 'c'",
			set:      testSet,
			prefix:   "c",
			expected: map[string]string{"cat": "animal", "dog": "animal"},
		},
		{
			name:     "Test case 3: Prefix 'd'",
			set:      testSet,
			prefix:   "d",
			expected: map[string]string{"dog": "animal"},
		},
		{
			name:     "Test case 4: Prefix 'e'",
			set:      testSet,
			prefix:   "e",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GetMapByPrefix(tc.set, tc.prefix)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
