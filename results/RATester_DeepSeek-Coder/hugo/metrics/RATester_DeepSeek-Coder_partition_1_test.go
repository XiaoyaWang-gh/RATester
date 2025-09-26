package metrics

import (
	"fmt"
	"testing"
)

func TestPartition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		d, scale, expected int
	}

	tests := []test{
		{10, 3, 9},
		{100, 10, 100},
		{1000, 100, 1000},
		{10000, 1000, 10000},
		{100000, 10000, 100000},
	}

	for _, test := range tests {
		result := partition(test.d, test.scale)
		if result != test.expected {
			t.Errorf("partition(%d, %d) = %d; want %d", test.d, test.scale, result, test.expected)
		}
	}
}
