package memory

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_3(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected []byte
		err      error
	}

	tests := []test{
		{
			name:     "Test case 1: Get existing key",
			key:      "key1",
			expected: []byte("value1"),
			err:      nil,
		},
		{
			name:     "Test case 2: Get non-existing key",
			key:      "key2",
			expected: nil,
			err:      nil,
		},
		{
			name:     "Test case 3: Get key with expired value",
			key:      "key3",
			expected: nil,
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			storage := &Storage{
				db: map[string]entry{
					"key1": {data: []byte("value1"), expiry: 0},
					"key3": {data: []byte("value3"), expiry: 1},
				},
			}

			value, err := storage.Get(test.key)

			if !reflect.DeepEqual(value, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, value)
			}

			if !errors.Is(err, test.err) {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
		})
	}
}
