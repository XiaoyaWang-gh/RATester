package metrics

import (
	"fmt"
	"testing"
)

func TestHowSimilarStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		a, b     string
		expected int
	}

	tests := []test{
		{"", "", 100},
		{"hello world", "hello world", 100},
		{"hello world", "hello", 25},
		{"hello", "hello world", 25},
		{"hello world", "world hello", 50},
		{"hello world", "hello hello", 50},
		{"hello world", "world world", 50},
		{"hello world", "hello world hello", 75},
		{"hello world", "world hello world", 75},
		{"hello world", "hello world world", 75},
		{"hello world", "world hello hello", 75},
		{"hello world", "hello hello world", 75},
		{"hello world", "world world hello", 75},
		{"hello world", "hello world hello world", 100},
		{"hello world", "world hello world hello", 100},
		{"hello world", "hello world world hello", 100},
		{"hello world", "world hello hello world", 100},
		{"hello world", "hello world hello hello", 100},
		{"hello world", "world world world hello", 100},
	}

	for _, test := range tests {
		result := howSimilarStrings(test.a, test.b)
		if result != test.expected {
			t.Errorf("howSimilarStrings(%q, %q) = %v; want %v", test.a, test.b, result, test.expected)
		}
	}
}
