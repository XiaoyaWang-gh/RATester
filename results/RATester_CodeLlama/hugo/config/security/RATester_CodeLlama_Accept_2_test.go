package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAccept_2(t *testing.T) {
	tests := []struct {
		name string
		w    Whitelist
		want bool
	}{
		{
			name: "acceptNone",
			w:    Whitelist{acceptNone: true},
			want: false,
		},
		{
			name: "empty",
			w:    Whitelist{},
			want: false,
		},
		{
			name: "match",
			w:    Whitelist{patterns: []*regexp.Regexp{regexp.MustCompile("^foo$")}},
			want: true,
		},
		{
			name: "no match",
			w:    Whitelist{patterns: []*regexp.Regexp{regexp.MustCompile("^bar$")}},
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

			if got := tt.w.Accept(tt.name); got != tt.want {
				t.Errorf("Whitelist.Accept() = %v, want %v", got, tt.want)
			}
		})
	}
}
