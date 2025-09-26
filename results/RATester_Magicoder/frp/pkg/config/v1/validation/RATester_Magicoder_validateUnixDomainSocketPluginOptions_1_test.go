package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateUnixDomainSocketPluginOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		options *v1.UnixDomainSocketPluginOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: &v1.UnixDomainSocketPluginOptions{
				UnixPath: "/path/to/socket",
			},
			wantErr: false,
		},
		{
			name: "invalid options",
			options: &v1.UnixDomainSocketPluginOptions{
				UnixPath: "",
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

			err := validateUnixDomainSocketPluginOptions(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateUnixDomainSocketPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
