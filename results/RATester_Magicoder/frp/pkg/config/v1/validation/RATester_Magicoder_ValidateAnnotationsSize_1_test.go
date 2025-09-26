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
			name: "test case 1",
			annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
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

			if err := ValidateAnnotationsSize(tt.annotations); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAnnotationsSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
