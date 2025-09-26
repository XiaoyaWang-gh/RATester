package tailscale

import (
	"fmt"
	"testing"
)

func TestIsTailscaleDomain_1(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{domain: "test.ts.net"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{domain: "test.ts.com"},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{domain: "test.net"},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{domain: "test.ts.net.au"},
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

			if got := isTailscaleDomain(tt.args.domain); got != tt.want {
				t.Errorf("isTailscaleDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
