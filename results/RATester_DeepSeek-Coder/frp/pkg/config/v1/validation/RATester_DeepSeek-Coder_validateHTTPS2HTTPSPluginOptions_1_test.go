package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTPS2HTTPSPluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.HTTPS2HTTPSPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &v1.HTTPS2HTTPSPluginOptions{
					LocalAddr: "127.0.0.1:8080",
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				c: &v1.HTTPS2HTTPSPluginOptions{
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

			if err := validateHTTPS2HTTPSPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateHTTPS2HTTPSPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
