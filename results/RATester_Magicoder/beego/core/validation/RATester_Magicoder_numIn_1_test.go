package validation

import (
	"fmt"
	"testing"
)

func TestnumIn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		name     string
		expected int
		err      error
	}{
		{"test1", 1, nil},
		{"test2", 2, nil},
		{"test3", 3, nil},
		{"test4", 4, nil},
		{"test5", 5, nil},
	}

	for _, test := range tests {
		num, err := numIn(test.name)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error %v, but got %v", test.err, err)
		}
		if num != test.expected {
			t.Errorf("Expected num %d, but got %d", test.expected, num)
		}
	}
}
