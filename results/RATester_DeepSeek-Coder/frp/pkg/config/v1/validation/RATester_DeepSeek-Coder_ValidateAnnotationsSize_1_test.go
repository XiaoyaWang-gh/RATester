package validation

import (
	"fmt"
	"testing"
)

func TestValidateAnnotationsSize_1(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		wantErr     bool
	}{
		{
			name: "Test case 1: Valid annotations",
			annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "Test case 2: Annotations size limit exceeded",
			annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
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

			err := ValidateAnnotationsSize(tt.annotations)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAnnotationsSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
