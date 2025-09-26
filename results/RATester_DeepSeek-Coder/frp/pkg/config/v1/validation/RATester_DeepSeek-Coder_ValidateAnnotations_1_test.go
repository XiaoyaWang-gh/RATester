package validation

import (
	"fmt"
	"testing"
)

func TestValidateAnnotations_1(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		wantErr     bool
	}{
		{
			name:        "valid annotations",
			annotations: map[string]string{"key1": "value1", "key2": "value2"},
			wantErr:     false,
		},
		{
			name:        "empty annotations",
			annotations: map[string]string{},
			wantErr:     false,
		},
		{
			name:        "invalid key",
			annotations: map[string]string{"invalid/key": "value1"},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := ValidateAnnotations(tt.annotations)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAnnotations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
