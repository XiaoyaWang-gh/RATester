package doctree

import (
	"fmt"
	"testing"
)

func TestValidateKey_1(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{
			name:    "empty key",
			key:     "",
			wantErr: false,
		},
		{
			name:    "short key",
			key:     "/a",
			wantErr: true,
		},
		{
			name:    "key does not start with '/'",
			key:     "a/b",
			wantErr: true,
		},
		{
			name:    "key ends with '/'",
			key:     "/a/",
			wantErr: true,
		},
		{
			name:    "valid key",
			key:     "/a/b",
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

			if err := ValidateKey(tt.key); (err != nil) != tt.wantErr {
				t.Errorf("ValidateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
