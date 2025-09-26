package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateProxyBaseConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.ProxyBaseConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				c: &v1.ProxyBaseConfig{
					Annotations: map[string]string{
						"test": "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				c: &v1.ProxyBaseConfig{
					Annotations: map[string]string{
						"test": "test",
					},
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

			if err := validateProxyBaseConfigForServer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateProxyBaseConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
