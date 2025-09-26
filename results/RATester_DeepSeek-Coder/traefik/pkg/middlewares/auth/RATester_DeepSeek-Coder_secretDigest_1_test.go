package auth

import (
	"fmt"
	"testing"
)

func TestSecretDigest_1(t *testing.T) {
	d := &digestAuth{
		users: map[string]string{
			"testuser:testrealm": "testsecret",
		},
	}

	type args struct {
		user  string
		realm string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with existing user and realm",
			args: args{
				user:  "testuser",
				realm: "testrealm",
			},
			want: "testsecret",
		},
		{
			name: "Test with non-existing user and realm",
			args: args{
				user:  "nonexistinguser",
				realm: "nonexistingrealm",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := d.secretDigest(tt.args.user, tt.args.realm); got != tt.want {
				t.Errorf("secretDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}
