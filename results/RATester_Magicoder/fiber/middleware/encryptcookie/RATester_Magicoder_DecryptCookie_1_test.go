package encryptcookie

import (
	"fmt"
	"testing"
)

func TestDecryptCookie_1(t *testing.T) {
	key := "YELLOW SUBMARINE"
	value := "encrypted value"

	tests := []struct {
		name    string
		value   string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "valid",
			value:   value,
			key:     key,
			want:    "plaintext",
			wantErr: false,
		},
		{
			name:    "invalid key",
			value:   value,
			key:     "invalid key",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid value",
			value:   "invalid value",
			key:     key,
			want:    "",
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

			got, err := DecryptCookie(tt.value, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
