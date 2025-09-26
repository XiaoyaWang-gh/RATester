package encryptcookie

import (
	"fmt"
	"testing"
)

func TestGenerateKey_1(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{
			name:    "Test case 1: Valid key length 16",
			length:  16,
			wantErr: false,
		},
		{
			name:    "Test case 2: Valid key length 24",
			length:  24,
			wantErr: false,
		},
		{
			name:    "Test case 3: Valid key length 32",
			length:  32,
			wantErr: false,
		},
		{
			name:    "Test case 4: Invalid key length 0",
			length:  0,
			wantErr: true,
		},
		{
			name:    "Test case 5: Invalid key length -1",
			length:  -1,
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

			defer func() {
				if r := recover(); (r != nil) != tt.wantErr {
					t.Errorf("GenerateKey() = %v, wantErr %v", r, tt.wantErr)
				}
			}()

			got := GenerateKey(tt.length)
			if tt.wantErr {
				t.Errorf("GenerateKey() = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
