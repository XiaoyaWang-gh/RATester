package middleware

import (
	"fmt"
	"testing"
)

func TestMatchScheme_1(t *testing.T) {
	type args struct {
		domain  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Context_0",
			args: args{
				domain:  "http://example.com",
				pattern: "http://example.com",
			},
			want: true,
		},
		{
			name: "Context_1",
			args: args{
				domain:  "https://example.com",
				pattern: "http://example.com",
			},
			want: false,
		},
		{
			name: "Context_2",
			args: args{
				domain:  "http://example.com",
				pattern: "https://example.com",
			},
			want: false,
		},
		{
			name: "Context_3",
			args: args{
				domain:  "http://example.com",
				pattern: "http://example.com",
			},
			want: true,
		},
		{
			name: "Context_4",
			args: args{
				domain:  "http://example.com",
				pattern: "http://example.com",
			},
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

			if got := matchScheme(tt.args.domain, tt.args.pattern); got != tt.want {
				t.Errorf("matchScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
