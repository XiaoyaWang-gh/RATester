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
		user  string
		realm string
		want  string
	}{
		{"user1", "realm1", "password1"},
		{"user2", "realm2", "password2"},
		{"user3", "realm3", ""},
	}

	for _, tt := range tests {
		t.Run(tt.user, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := b.secretBasic(tt.user, tt.realm); got != tt.want {
				t.Errorf("secretBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
