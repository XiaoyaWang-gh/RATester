package template

import (
	"fmt"
	"testing"
)

func TestIsCSSSpace_1(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{
			name: "test case 1",
			b:    '\t',
			want: true,
		},
		{
			name: "test case 2",
			b:    '\n',
			want: true,
		},
		{
			name: "test case 3",
			b:    '\f',
			want: true,
		},
		{
			name: "test case 4",
			b:    '\r',
			want: true,
		},
		{
			name: "test case 5",
			b:    ' ',
			want: true,
		},
		{
			name: "test case 6",
			b:    'a',
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isCSSSpace(tt.b); got != tt.want {
				t.Errorf("isCSSSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
