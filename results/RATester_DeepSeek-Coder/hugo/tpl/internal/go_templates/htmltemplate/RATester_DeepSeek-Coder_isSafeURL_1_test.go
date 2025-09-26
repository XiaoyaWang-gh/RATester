package template

import (
	"fmt"
	"testing"
)

func TestIsSafeURL_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "Test case 1",
			arg:  "http://example.com",
			want: true,
		},
		{
			name: "Test case 2",
			arg:  "https://example.com",
			want: true,
		},
		{
			name: "Test case 3",
			arg:  "mailto:example@example.com",
			want: true,
		},
		{
			name: "Test case 4",
			arg:  "ftp://example.com",
			want: false,
		},
		{
			name: "Test case 5",
			arg:  "example.com",
			want: false,
		},
		{
			name: "Test case 6",
			arg:  "http:/example.com",
			want: false,
		},
		{
			name: "Test case 7",
			arg:  "https:/example.com",
			want: false,
		},
		{
			name: "Test case 8",
			arg:  "mailto:/example.com",
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

			if got := isSafeURL(tt.arg); got != tt.want {
				t.Errorf("isSafeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
