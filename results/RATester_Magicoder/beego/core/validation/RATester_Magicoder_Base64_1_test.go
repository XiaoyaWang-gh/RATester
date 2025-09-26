package validation

import (
	"fmt"
	"testing"
)

func TestBase64_1(t *testing.T) {
	v := Validation{}

	tests := []struct {
		name    string
		obj     interface{}
		key     string
		wantErr bool
	}{
		{
			name:    "valid base64",
			obj:     "SGVsbG8gd29ybGQ=",
			key:     "fieldName",
			wantErr: false,
		},
		{
			name:    "invalid base64",
			obj:     "invalid base64",
			key:     "fieldName",
			wantErr: true,
		},
		{
			name:    "empty string",
			obj:     "",
			key:     "fieldName",
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

			got := v.Base64(tt.obj, tt.key)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("Base64() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
		})
	}
}
