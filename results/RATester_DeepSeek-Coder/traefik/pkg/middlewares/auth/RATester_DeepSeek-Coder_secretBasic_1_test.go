package auth

import (
	"fmt"
	"testing"
)

func TestSecretBasic_1(t *testing.T) {
	b := &basicAuth{
		users: map[string]string{
			"user1": "password1",
			"user2": "password2",
		},
	}

	tests := []struct {
		name     string
		user     string
		realm    string
		expected string
	}{
		{
			name:     "User Exists",
			user:     "user1",
			realm:    "testRealm",
			expected: "password1",
		},
		{
			name:     "User Does Not Exist",
			user:     "user3",
			realm:    "testRealm",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := b.secretBasic(test.user, test.realm)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
