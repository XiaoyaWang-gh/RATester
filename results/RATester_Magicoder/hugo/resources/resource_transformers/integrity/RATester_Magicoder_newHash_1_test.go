package integrity

import (
	"fmt"
	"testing"
)

func TestnewHash_1(t *testing.T) {
	tests := []struct {
		name    string
		algo    string
		wantErr bool
	}{
		{"md5", "md5", false},
		{"sha256", "sha256", false},
		{"sha384", "sha384", true},
		{"sha512", "sha512", false},
		{"unsupported", "unsupported", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newHash(tt.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("newHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
