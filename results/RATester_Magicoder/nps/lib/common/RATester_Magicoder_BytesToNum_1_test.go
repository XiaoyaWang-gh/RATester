package common

import (
	"fmt"
	"testing"
)

func TestBytesToNum_1(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want int
	}{
		{
			name: "Test Case 1",
			b:    []byte{1, 2, 3},
			want: 123,
		},
		{
			name: "Test Case 2",
			b:    []byte{4, 5, 6},
			want: 456,
		},
		{
			name: "Test Case 3",
			b:    []byte{7, 8, 9},
			want: 789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BytesToNum(tt.b); got != tt.want {
				t.Errorf("BytesToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
