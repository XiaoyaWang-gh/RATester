package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetWriteStr_1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected []byte
	}{
		{
			name:     "test case 1",
			input:    []string{"123", "456"},
			expected: []byte{0x31, 0x32, 0x33, 0x03, 0x34, 0x35, 0x36},
		},
		{
			name:     "test case 2",
			input:    []string{"123", "456", "789"},
			expected: []byte{0x31, 0x32, 0x33, 0x03, 0x34, 0x35, 0x36, 0x03, 0x37, 0x38, 0x39},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetWriteStr(tc.input...); !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("GetWriteStr() = %v, want %v", got, tc.expected)
			}
		})
	}
}
