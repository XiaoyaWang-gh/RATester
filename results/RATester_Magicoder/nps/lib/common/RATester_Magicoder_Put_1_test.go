package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPut_1(t *testing.T) {
	pool := &copyBufferPool{}
	pool.New()

	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "valid input",
			input:    make([]byte, PoolSizeCopy),
			expected: make([]byte, PoolSizeCopy),
		},
		{
			name:     "invalid input",
			input:    make([]byte, PoolSizeCopy+1),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pool.Put(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, tc.input)
			}
		})
	}
}
