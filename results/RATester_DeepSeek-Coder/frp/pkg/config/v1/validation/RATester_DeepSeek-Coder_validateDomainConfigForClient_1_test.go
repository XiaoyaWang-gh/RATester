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
			name: "both subdomain and custom domains are empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{},
					SubDomain:     "",
				},
			},
			wantErr: true,
		},
		{
			name: "only subdomain is not empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{},
					SubDomain:     "test",
				},
			},
			wantErr: false,
		},
		{
			name: "only custom domains are not empty",
			args: args{
				c: &v1.DomainConfig{
					CustomDomains: []string{"test.com"},
					SubDomain:     "",
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

			err := validateDomainConfigForClient(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDomainConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
