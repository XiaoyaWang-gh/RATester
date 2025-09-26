package vhost

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestParseBasicAuth_1(t *testing.T) {
	tests := []struct {
		name     string
		auth     string
		expected struct {
			username string
			password string
			ok       bool
		}
	}{
		{
			name: "valid auth",
			auth: "Basic " + base64.StdEncoding.EncodeToString([]byte("username:password")),
			expected: struct {
				username string
				password string
				ok       bool
			}{
				username: "username",
				password: "password",
				ok:       true,
			},
		},
		{
			name: "invalid auth",
			auth: "Basic " + base64.StdEncoding.EncodeToString([]byte("username")),
			expected: struct {
				username string
				password string
				ok       bool
			}{},
		},
		{
			name: "empty auth",
			auth: "",
			expected: struct {
				username string
				password string
				ok       bool
			}{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			username, password, ok := parseBasicAuth(test.auth)
			if username != test.expected.username || password != test.expected.password || ok != test.expected.ok {
				t.Errorf("Expected (%s, %s, %t), got (%s, %s, %t)", test.expected.username, test.expected.password, test.expected.ok, username, password, ok)
			}
		})
	}
}
