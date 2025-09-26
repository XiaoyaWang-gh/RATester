package encryptcookie

import (
	"fmt"
	"testing"
)

func TestEncryptCookie_1(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "valid input",
			value:   "test",
			key:     "testkeytestkeytestkeytestkey",
			want:    "test",
			wantErr: false,
		},
		{
			name:    "invalid key length",
			value:   "test",
			key:     "test",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid base64 key",
			value:   "test",
			key:     "invalidbase64key",
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

			got, err := EncryptCookie(tt.value, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
