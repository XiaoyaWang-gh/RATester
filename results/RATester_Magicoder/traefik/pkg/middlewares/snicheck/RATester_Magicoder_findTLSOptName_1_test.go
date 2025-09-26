package snicheck

import (
	"fmt"
	"testing"
)

func TestFindTLSOptName_1(t *testing.T) {
	type args struct {
		tlsOptionsForHost map[string]string
		host              string
		fqdn              bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions1"},
				host:              "example.com",
				fqdn:              false,
			},
			want: "tlsOptions1",
		},
		{
			name: "Test Case 2",
			args: args{
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions2"},
				host:              "example.com.",
				fqdn:              true,
			},
			want: "tlsOptions2",
		},
		{
			name: "Test Case 3",
			args: args{
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions3"},
				host:              "example.com",
				fqdn:              true,
			},
			want: "tlsOptions3",
		},
		{
			name: "Test Case 4",
			args: args{
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions4"},
				host:              "example.com",
				fqdn:              false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findTLSOptName(tt.args.tlsOptionsForHost, tt.args.host, tt.args.fqdn); got != tt.want {
				t.Errorf("findTLSOptName() = %v, want %v", got, tt.want)
			}
		})
	}
}
