package pagemeta

import (
	"fmt"
	"testing"
)

func TestValidate_1(t *testing.T) {
	tests := []struct {
		name    string
		rc      ResourceConfig
		wantErr bool
	}{
		{
			name: "valid",
			rc: ResourceConfig{
				Path: "valid_path",
			},
			wantErr: false,
		},
		{
			name: "invalid_path",
			rc: ResourceConfig{
				Path: "",
			},
			wantErr: true,
		},
		{
			name: "invalid_markup",
			rc: ResourceConfig{
				Path: "valid_path",
				Content: Source{
					Markup: "markdown",
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

			if err := tt.rc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ResourceConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
