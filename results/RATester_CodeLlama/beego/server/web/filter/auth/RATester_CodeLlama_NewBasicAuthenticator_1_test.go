package auth

import (
	"fmt"
	"testing"
)

func TestNewBasicAuthenticator_1(t *testing.T) {
	type args struct {
		secrets SecretProvider
		Realm   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				secrets: func(user, pass string) bool {
					return true
				},
				Realm: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewBasicAuthenticator(tt.args.secrets, tt.args.Realm); got == nil {
				t.Errorf("NewBasicAuthenticator() = %v, want %v", got, tt.name)
			}
		})
	}
}
