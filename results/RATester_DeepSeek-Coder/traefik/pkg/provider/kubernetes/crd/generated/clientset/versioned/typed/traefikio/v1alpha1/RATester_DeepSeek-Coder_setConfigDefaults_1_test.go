package v1alpha1

import (
	"fmt"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestSetConfigDefaults_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *rest.Config
		wantErr bool
	}{
		{
			name: "Test with empty config",
			config: &rest.Config{
				Host: "",
			},
			wantErr: false,
		},
		{
			name: "Test with non-empty config",
			config: &rest.Config{
				Host: "http://localhost:8080",
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

			err := setConfigDefaults(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("setConfigDefaults() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.config.GroupVersion == nil {
				t.Errorf("setConfigDefaults() GroupVersion is nil")
			}
			if tt.config.APIPath != "/apis" {
				t.Errorf("setConfigDefaults() APIPath = %v, want %v", tt.config.APIPath, "/apis")
			}
			if tt.config.NegotiatedSerializer == nil {
				t.Errorf("setConfigDefaults() NegotiatedSerializer is nil")
			}
			if tt.config.UserAgent == "" {
				t.Errorf("setConfigDefaults() UserAgent is empty")
			}
		})
	}
}
