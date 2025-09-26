package metadecoders

import (
	"fmt"
	"testing"
)

func TestparseORGDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input  string
		output string
	}

	tests := []test{
		{"[2022-01-01]", "2022-01-01"},
		{"[2022-01-01]", "2022-01-01"},
		{"[2022-01-01]", "2022-01-01"},
		{"[2022-01-01]", "2022-01-01"},
		{"[2022-01-01]", "2022-01-01"},
	}

	for _, test := range tests {
		result := parseORGDate(test.input)
		if result != test.output {
			t.Errorf("Expected %s, got %s", test.output, result)
		}
	}
}
