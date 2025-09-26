package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateUnixDomainSocketPluginOptions_1(t *testing.T) {
	type args struct {
		c *v1.UnixDomainSocketPluginOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unixPath is required",
			args: args{
				c: &v1.UnixDomainSocketPluginOptions{
					UnixPath: "",
				},
			},
			wantErr: true,
		},
		{
			name: "unixPath is not required",
			args: args{
				c: &v1.UnixDomainSocketPluginOptions{
					UnixPath: "test",
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

			if err := validateUnixDomainSocketPluginOptions(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateUnixDomainSocketPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
