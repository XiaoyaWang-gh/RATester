package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExpandDefaultValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		values   []string
		defaults []string
		expected []string
	}

	tests := []test{
		{
			values:   []string{":default", "value1", "value2"},
			defaults: []string{"default1", "default2"},
			expected: []string{"default1", "default2", "value1", "value2"},
		},
		{
			values:   []string{"value1", "value2"},
			defaults: []string{"default1", "default2"},
			expected: []string{"value1", "value2"},
		},
		{
			values:   []string{":default"},
			defaults: []string{"default1", "default2"},
			expected: []string{"default1", "default2"},
		},
		{
			values:   []string{},
			defaults: []string{"default1", "default2"},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := expandDefaultValues(test.values, test.defaults)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expandDefaultValues(%v, %v) = %v, want %v", test.values, test.defaults, result, test.expected)
		}
	}
}
