package yaml

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		key      string
		expected []string
		err      error
	}{
		{
			name:     "success",
			key:      "key1",
			expected: []string{"val1", "val2"},
			err:      nil,
		},
		{
			name:     "key not found",
			key:      "key2",
			expected: nil,
			err:      errors.New("key not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cc := &ConfigContainer{
				data: map[string]interface{}{
					"key1": "val1;val2",
				},
			}

			actual, err := cc.Strings(tc.key)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
