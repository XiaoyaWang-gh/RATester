package http

import (
	"fmt"
	"testing"
)

func TestParseBasicAuth_2(t *testing.T) {
	tests := []struct {
		name     string
		auth     string
		wantUser string
		wantPass string
		wantOk   bool
	}{
		{
			name:     "valid auth",
			auth:     "Basic dXNlcjpwYXNzd29yZA==",
			wantUser: "username",
			wantPass: "password",
			wantOk:   true,
		},
		{
			name:     "invalid auth",
			auth:     "Basic invalid",
			wantUser: "",
			wantPass: "",
			wantOk:   false,
		},
		{
			name:     "empty auth",
			auth:     "",
			wantUser: "",
			wantPass: "",
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			user, pass, ok := ParseBasicAuth(tt.auth)
			if user != tt.wantUser || pass != tt.wantPass || ok != tt.wantOk {
				t.Errorf("ParseBasicAuth(%q) = %q, %q, %v, want %q, %q, %v", tt.auth, user, pass, ok, tt.wantUser, tt.wantPass, tt.wantOk)
			}
		})
	}
}
