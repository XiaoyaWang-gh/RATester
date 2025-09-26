package pageparser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestVal_1(t *testing.T) {
	type test struct {
		name     string
		item     Item
		source   []byte
		expected []byte
	}

	tests := []test{
		{
			name: "single segment",
			item: Item{
				low:  0,
				high: 10,
			},
			source:   []byte("0123456789"),
			expected: []byte("0123456789"),
		},
		{
			name: "multiple segments",
			item: Item{
				segments: []lowHigh{
					{Low: 0, High: 5},
					{Low: 5, High: 10},
				},
			},
			source:   []byte("0123456789"),
			expected: []byte("0123456789"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.item.Val(tt.source)
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
