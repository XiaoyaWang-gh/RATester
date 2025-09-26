package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbindataRead_1(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected []byte
		err      error
	}{
		{
			name:     "valid gzip data",
			data:     []byte{0x1F, 0x8B, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xFF, 0x24, 0x74, 0x4D, 0x4D, 0x01, 0x00, 0x3B},
			expected: []byte("Hello, World!"),
			err:      nil,
		},
		{
			name:     "invalid gzip data",
			data:     []byte{0x1F, 0x8B, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xFF, 0x24, 0x74, 0x4D, 0x4D, 0x01, 0x00, 0x3A},
			expected: nil,
			err:      fmt.Errorf("Read \"invalid gzip data\": gzip: invalid header"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := bindataRead(tc.data, tc.name)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected error: %v, Actual error: %v", tc.err, err)
			}
		})
	}
}
