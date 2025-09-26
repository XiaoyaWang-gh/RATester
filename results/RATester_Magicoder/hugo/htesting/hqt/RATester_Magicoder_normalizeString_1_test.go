package hqt

import (
	"fmt"
	"testing"
)

func TestnormalizeString_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "normal case",
			arg:  "Hello\r\nWorld",
			want: "Hello\nWorld",
		},
		{
			name: "multiple spaces",
			arg:  "Hello   World",
			want: "HelloWorld",
		},
		{
			name: "leading and trailing spaces",
			arg:  "   Hello   World   ",
			want: "HelloWorld",
		},
		{
			name: "multiple lines",
			arg:  "Hello\r\n   World   \r\n   Universe   ",
			want: "Hello\nWorld\nUniverse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := normalizeString(tt.arg); got != tt.want {
				t.Errorf("normalizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
