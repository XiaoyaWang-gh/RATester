package types

import (
	"fmt"
	"testing"
)

func TestMatchDomain_1(t *testing.T) {
	type args struct {
		domain     string
		certDomain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exact match",
			args: args{
				domain:     "example.com",
				certDomain: "example.com",
			},
			want: true,
		},
		{
			name: "wildcard match",
			args: args{
				domain:     "*.example.com",
				certDomain: "example.com",
			},
			want: true,
		},
		{
			name: "no match",
			args: args{
				domain:     "example.com",
				certDomain: "foo.example.com",
			},
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

			if got := MatchDomain(tt.args.domain, tt.args.certDomain); got != tt.want {
				t.Errorf("MatchDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
