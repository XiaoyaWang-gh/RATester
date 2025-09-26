package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRandomCreateBytes_1(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		alphabets []byte
	}{
		{
			name:      "Test with default alphabets",
			n:         10,
			alphabets: alphaNum,
		},
		{
			name:      "Test with custom alphabets",
			n:         15,
			alphabets: []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		},
		{
			name:      "Test with zero length alphabets",
			n:         5,
			alphabets: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := RandomCreateBytes(tt.n, tt.alphabets...)
			if len(got) != tt.n {
				t.Errorf("RandomCreateBytes() length = %v, want %v", len(got), tt.n)
				return
			}
			for _, b := range got {
				if len(tt.alphabets) > 0 && !bytes.Contains(tt.alphabets, []byte{b}) {
					t.Errorf("RandomCreateBytes() byte = %v, not in alphabets", b)
				}
			}
		})
	}
}
