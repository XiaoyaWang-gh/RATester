package metadecoders

import (
	"fmt"
	"testing"
)

func TestParseORGDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{"[2022-01-02 Mon]", "2022-01-02"},
		{"<2022-01-02 Mon>", "2022-01-02"},
		{"[2022-01-02]", "2022-01-02"},
		{"<2022-01-02>", "2022-01-02"},
		{"2022-01-02", "2022-01-02"},
		{"[2022-01-02 Mon 12:00]", "2022-01-02"},
		{"<2022-01-02 Mon 12:00>", "2022-01-02"},
	}

	for _, test := range tests {
		result := parseORGDate(test.input)
		if result != test.expected {
			t.Errorf("parseORGDate(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
