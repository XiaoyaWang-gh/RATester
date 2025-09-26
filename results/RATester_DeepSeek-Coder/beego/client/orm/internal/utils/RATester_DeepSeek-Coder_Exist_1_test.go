package utils

import (
	"fmt"
	"testing"
)

func TestExist_1(t *testing.T) {
	tests := []struct {
		name string
		f    StrTo
		want bool
	}{
		{
			name: "Test case 1",
			f:    StrTo("test"),
			want: true,
		},
		{
			name: "Test case 2",
			f:    StrTo(string(rune(0x1E))),
			want: false,
		},
		{
			name: "Test case 3",
			f:    StrTo(string(rune(0x01))),
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

			if got := tt.f.Exist(); got != tt.want {
				t.Errorf("StrTo.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
