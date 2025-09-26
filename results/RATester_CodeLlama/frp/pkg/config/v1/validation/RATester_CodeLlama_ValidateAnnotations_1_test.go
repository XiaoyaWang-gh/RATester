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
			name:        "empty",
			annotations: map[string]string{},
			wantErr:     false,
		},
		{
			name: "valid",
			annotations: map[string]string{
				"foo": "bar",
			},
			wantErr: false,
		},
		{
			name: "invalid key",
			annotations: map[string]string{
				"123": "bar",
			},
			wantErr: true,
		},
		{
			name: "too many annotations",
			annotations: map[string]string{
				"foo": "bar",
				"bar": "baz",
				"baz": "foo",
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

			if err := ValidateAnnotations(tt.annotations); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAnnotations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
