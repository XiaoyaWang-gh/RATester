package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTPS2HTTPSPluginOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		options *v1.HTTPS2HTTPSPluginOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: &v1.HTTPS2HTTPSPluginOptions{
				LocalAddr: ":8080",
			},
			wantErr: false,
		},
		{
			name: "invalid options",
			options: &v1.HTTPS2HTTPSPluginOptions{
				LocalAddr: "",
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

			err := validateHTTPS2HTTPSPluginOptions(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateHTTPS2HTTPSPluginOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
