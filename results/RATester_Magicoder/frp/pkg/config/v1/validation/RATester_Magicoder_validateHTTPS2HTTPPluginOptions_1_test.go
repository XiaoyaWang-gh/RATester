package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTPS2HTTPPluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.HTTPS2HTTPPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1: LocalAddr is empty",
			args: args{
				c: &v1.HTTPS2HTTPPluginOptions{
					LocalAddr: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 2: LocalAddr is not empty",
			args: args{
				c: &v1.HTTPS2HTTPPluginOptions{
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

			if err := validateHTTPS2HTTPPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateHTTPS2HTTPPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
