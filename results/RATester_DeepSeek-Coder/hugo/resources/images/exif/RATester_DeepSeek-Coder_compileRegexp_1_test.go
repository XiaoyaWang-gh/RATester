package exif

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestCompileRegexp_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *regexp.Regexp
		err      error
	}{
		{
			name:     "empty expression",
			input:    "",
			expected: nil,
			err:      nil,
		},
		{
			name:     "non-empty expression",
			input:    "test",
			expected: regexp.MustCompile("(?i)test"),
			err:      nil,
		},
		{
			name:     "expression with prefix",
			input:    "(?i)test",
			expected: regexp.MustCompile("(?i)test"),
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := compileRegexp(tt.input)
			if err != tt.err {
				t.Errorf("Expected error %v, got %v", tt.err, err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
