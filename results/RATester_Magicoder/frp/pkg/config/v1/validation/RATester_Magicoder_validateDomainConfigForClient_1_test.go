package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateDomainConfigForClient_1(t *testing.T) {
	type args struct {
		c *v1.DomainConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "both empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{},
					SubDomain:     "",
				},
			},
			wantErr: true,
		},
		{
			name: "subdomain empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{"example.com"},
					SubDomain:     "",
				},
			},
			wantErr: false,
		},
		{
			name: "custom domains empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{},
					SubDomain:     "example.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := validateDomainConfigForClient(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateDomainConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
