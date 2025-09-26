package session

import (
	"fmt"
	"testing"
)

func TestCfgDomain_1(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{domain: "example.com"},
			want: func(config *ManagerConfig) {
				config.Domain = "example.com"
			},
		},
		{
			name: "Test case 2",
			args: args{domain: "example.net"},
			want: func(config *ManagerConfig) {
				config.Domain = "example.net"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := CfgDomain(tt.args.domain)
			config := &ManagerConfig{}
			got(config)
			if config.Domain != tt.args.domain {
				t.Errorf("CfgDomain() = %v, want %v", config.Domain, tt.args.domain)
			}
		})
	}
}
