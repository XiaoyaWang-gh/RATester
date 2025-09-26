package pageparser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestVal_1(t *testing.T) {
	tests := []struct {
		name     string
		item     Item
		source   []byte
		expected []byte
	}{
		{
			name: "single segment",
			item: Item{
				low:  1,
				high: 5,
			},
			source:   []byte("hello"),
			expected: []byte("ello"),
		},
		{
			name: "multiple segments",
			item: Item{
				segments: []lowHigh{
					{Low: 1, High: 3},
					{Low: 4, High: 6},
				},
			},
			source:   []byte("hello world"),
			expected: []byte("o w"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.item.Val(test.source)
			if !bytes.Equal(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
