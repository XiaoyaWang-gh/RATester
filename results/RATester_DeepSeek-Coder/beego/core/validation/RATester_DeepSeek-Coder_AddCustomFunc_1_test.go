package validation

import (
	"fmt"
	"testing"
)

func TestAddCustomFunc_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid function name",
			input:   "testFunc",
			wantErr: false,
		},
		{
			name:    "invalid function name",
			input:   "",
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

			err := AddCustomFunc(tt.input, func(v *Validation, obj interface{}, key string) {
				// custom function implementation
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("AddCustomFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
