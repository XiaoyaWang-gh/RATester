package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		name     string
		key      string
		s        []string
		expected []interface{}
		err      error
	}

	tests := []test{
		{
			name:     "validFunc",
			key:      "key1",
			s:        []string{"1", "2", "3"},
			expected: []interface{}{1, 2, 3, "key1"},
			err:      nil,
		},
		{
			name:     "invalidFunc",
			key:      "key2",
			s:        []string{"1", "2", "3"},
			expected: nil,
			err:      fmt.Errorf("doesn't exists %s valid function", "invalidFunc"),
		},
	}

	for _, test := range tests {
		result, err := trim(test.name, test.key, test.s)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error %v, got %v", test.err, err)
		}
	}
}
