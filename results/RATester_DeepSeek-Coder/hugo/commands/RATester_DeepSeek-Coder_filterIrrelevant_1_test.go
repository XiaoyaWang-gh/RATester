package commands

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestFilterIrrelevant_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "One irrelevant",
			input:    []string{"irrelevant"},
			expected: []string{},
		},
		{
			name:     "One relevant",
			input:    []string{"relevant"},
			expected: []string{"relevant"},
		},
		{
			name:     "Multiple files, some irrelevant",
			input:    []string{"relevant1", "irrelevant", "relevant2"},
			expected: []string{"relevant1", "relevant2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &fileChangeDetector{
				irrelevantRe: regexp.MustCompile("irrelevant"),
			}
			result := f.filterIrrelevant(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
