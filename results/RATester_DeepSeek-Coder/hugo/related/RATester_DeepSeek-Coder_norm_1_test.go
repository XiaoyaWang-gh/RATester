package related

import (
	"fmt"
	"testing"
)

func TestNorm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		num, min, max, expected int
	}

	tests := []test{
		{10, 0, 100, 10},
		{50, 0, 100, 50},
		{100, 0, 100, 100},
		{0, 0, 100, 0},
		{-10, 0, 100, 0},
		{110, 0, 100, 100},
	}

	for _, test := range tests {
		result := norm(test.num, test.min, test.max)
		if result != test.expected {
			t.Errorf("norm(%d, %d, %d) = %d; want %d", test.num, test.min, test.max, result, test.expected)
		}
	}
}
