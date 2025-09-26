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
			name: "Test case 1",
			args: args{
				c: &v1.ProxyBaseConfig{
					Name: "test",
					Type: "tcp",
					Annotations: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				c: &v1.ProxyBaseConfig{
					Name: "test",
					Type: "tcp",
					Annotations: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
				},
			},
			wantErr: true,
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
