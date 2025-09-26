package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetArgs_2(t *testing.T) {
	testCases := []struct {
		name     string
		args     []interface{}
		expected []interface{}
	}{
		{
			name:     "Test case 1",
			args:     []interface{}{1, "test", 3.14},
			expected: []interface{}{1, "test", 3.14},
		},
		{
			name:     "Test case 2",
			args:     []interface{}{"hello", "world"},
			expected: []interface{}{"hello", "world"},
		},
		{
			name:     "Test case 3",
			args:     []interface{}{100, 200, 300},
			expected: []interface{}{100, 200, 300},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := rawSet{}
			result := o.SetArgs(tc.args...)
			if !reflect.DeepEqual(result.(*rawSet).args, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result.(*rawSet).args)
			}
		})
	}
}
