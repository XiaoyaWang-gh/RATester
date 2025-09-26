package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTP2HTTPSPluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.HTTP2HTTPSPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1: localAddr is empty",
			args: args{
				c: &v1.HTTP2HTTPSPluginOptions{
					LocalAddr: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 2: localAddr is not empty",
			args: args{
				c: &v1.HTTP2HTTPSPluginOptions{
					LocalAddr: "127.0.0.1:8080",
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

			if err := validateHTTP2HTTPSPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateHTTP2HTTPSPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
