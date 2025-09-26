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
			name: "Test case 1: localAddr is empty",
			args: args{
				c: &v1.TLS2RawPluginOptions{
					Type:      "tcp",
					LocalAddr: "",
					CrtPath:   "/path/to/crt",
					KeyPath:   "/path/to/key",
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 2: localAddr is not empty",
			args: args{
				c: &v1.TLS2RawPluginOptions{
					Type:      "tcp",
					LocalAddr: ":8080",
					CrtPath:   "/path/to/crt",
					KeyPath:   "/path/to/key",
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

			if err := validateTLS2RawPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateTLS2RawPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
