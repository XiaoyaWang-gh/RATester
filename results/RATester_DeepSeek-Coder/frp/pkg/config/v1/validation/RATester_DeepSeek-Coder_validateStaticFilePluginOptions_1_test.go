package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateStaticFilePluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.StaticFilePluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				c: &v1.StaticFilePluginOptions{
					LocalPath: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				c: &v1.StaticFilePluginOptions{
					LocalPath: "",
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

			if err := validateStaticFilePluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateStaticFilePluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
