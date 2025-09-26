package media

import (
	"fmt"
	"testing"
)

func TestIsHTML_1(t *testing.T) {
	tests := []struct {
		name string
		m    Type
		want bool
	}{
		{
			name: "Test case 1",
			m:    Type{SubType: "html"},
			want: true,
		},
		{
			name: "Test case 2",
			m:    Type{SubType: "xml"},
			want: false,
		},
		{
			name: "Test case 3",
			m:    Type{SubType: "HTMLType.SubType"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.IsHTML(); got != tt.want {
				t.Errorf("IsHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
