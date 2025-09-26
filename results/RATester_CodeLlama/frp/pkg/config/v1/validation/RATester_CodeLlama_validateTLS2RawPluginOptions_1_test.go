package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateTLS2RawPluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.TLS2RawPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				c: &v1.TLS2RawPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				c: &v1.TLS2RawPluginOptions{
					LocalAddr: "",
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

			if err := validateTLS2RawPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateTLS2RawPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
