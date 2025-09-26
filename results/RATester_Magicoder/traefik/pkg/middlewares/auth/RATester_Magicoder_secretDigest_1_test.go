package auth

import (
	"fmt"
	"testing"
)

func TestSecretDigest_1(t *testing.T) {
	d := &digestAuth{
		users: map[string]string{
			"user1:realm1": "secret1",
			"user2:realm2": "secret2",
		},
	}

	tests := []struct {
		user  string
		realm string
		want  string
	}{
		{"user1", "realm1", "secret1"},
		{"user2", "realm2", "secret2"},
		{"user3", "realm3", ""},
	}

	for _, tt := range tests {
		t.Run(tt.user+":"+tt.realm, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := d.secretDigest(tt.user, tt.realm); got != tt.want {
				t.Errorf("secretDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}
