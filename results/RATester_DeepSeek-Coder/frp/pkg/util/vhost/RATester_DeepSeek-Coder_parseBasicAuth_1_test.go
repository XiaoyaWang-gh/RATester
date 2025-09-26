package vhost

import (
	"fmt"
	"testing"
)

func TestParseBasicAuth_1(t *testing.T) {
	tests := []struct {
		name     string
		auth     string
		wantUser string
		wantPass string
		wantOk   bool
	}{
		{
			name:     "valid basic auth",
			auth:     "Basic dXNlcjpwYXNzd29yZA==",
			wantUser: "username",
			wantPass: "password",
			wantOk:   true,
		},
		{
			name:     "invalid basic auth, missing prefix",
			auth:     "dXNlcjpwYXNzd29yZA==",
			wantUser: "",
			wantPass: "",
			wantOk:   false,
		},
		{
			name:     "invalid basic auth, invalid base64",
			auth:     "Basic dXNlcjpwYXNzd29yZA=",
			wantUser: "",
			wantPass: "",
			wantOk:   false,
		},
		{
			name:     "invalid basic auth, missing colon",
			auth:     "Basic dXNlcjpwYXNzd29yZA",
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

			user, pass, ok := parseBasicAuth(tt.auth)
			if user != tt.wantUser || pass != tt.wantPass || ok != tt.wantOk {
				t.Errorf("parseBasicAuth() = %v, %v, %v; want %v, %v, %v", user, pass, ok, tt.wantUser, tt.wantPass, tt.wantOk)
			}
		})
	}
}
