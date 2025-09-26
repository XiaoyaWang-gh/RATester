package test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMustAsset_1(t *testing.T) {
	testCases := []struct {
		name     string
		asset    string
		expected []byte
	}{
		{
			name:     "Test case 1",
			asset:    "asset1.txt",
			expected: []byte("This is asset1"),
		},
		{
			name:     "Test case 2",
			asset:    "asset2.txt",
			expected: []byte("This is asset2"),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := MustAsset(tc.asset)
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
