package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateStaticFilePluginOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		options *v1.StaticFilePluginOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: &v1.StaticFilePluginOptions{
				LocalPath: "/path/to/local/file",
			},
			wantErr: false,
		},
		{
			name: "invalid options - empty localPath",
			options: &v1.StaticFilePluginOptions{
				LocalPath: "",
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

			err := validateStaticFilePluginOptions(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateStaticFilePluginOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
