package encryptcookie

import (
	"fmt"
	"testing"
)

func TestEncryptCookie_1(t *testing.T) {
	type args struct {
		value string
		key   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				value: "test_value",
				key:   "test_key",
			},
			want:    "encrypted_value",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				value: "another_test_value",
				key:   "another_test_key",
			},
			want:    "another_encrypted_value",
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := EncryptCookie(tt.args.value, tt.args.key)
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
